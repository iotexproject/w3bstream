package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/output/chain"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/test/contract/powerc20"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
	"github.com/pkg/errors"
)

type Processor struct {
	vmHandler                 *vm.Handler
	projectManager            *project.Manager
	outputFactory             *output.Factory
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
	ps                        *p2p.PubSubs
}

func (r *Processor) handleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Task == nil {
		return
	}
	tid := d.Task.ID
	ms := d.Task.Messages
	slog.Debug("get new task", "task_id", tid)
	r.reportSuccess(tid, types.TaskStateFetched, "", topic)

	project, err := r.projectManager.Get(ms[0].ProjectID)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail(tid, err, topic)
		return
	}

	r.reportSuccess(tid, types.TaskStateProving, "", topic)
	res, err := r.vmHandler.Handle(ms, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail(tid, err, topic)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess(tid, types.TaskStateProved, string(res), topic)

	// output proof
	outCfg := r.buildProjectOutputConfig(project.Config)
	outputter, err := r.outputFactory.NewOutputter(outCfg)
	if err != nil {
		err = errors.Wrap(err, "fail to init outputter")
		slog.Error(err.Error())
		r.reportFail(tid, err, topic)
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", outputter))

	r.reportSuccess(tid, types.TaskStateOutputting, "output proof", topic)
	var outRes interface{}
	if project.ID == 20000 {
		outRes, err = r.writePowerc20(ms[0].Data, "https://babel-api.testnet.iotex.io", outCfg.ContractAddress, res)
		if err != nil {
			slog.Error(err.Error())
			r.reportFail(tid, err, topic)
			return
		}
	} else {
		outRes, err = outputter.Output(res)
		if err != nil {
			slog.Error(err.Error())
			r.reportFail(tid, err, topic)
			return
		}
	}
	r.reportSuccess(tid, types.TaskStateOutputted, fmt.Sprintf("output result: %+v", outRes), topic)
	slog.Debug("output success", "result", outRes)
}

// TODO merge powerc20 to output
func (r *Processor) writePowerc20(data, chainEndpoint, contractAddress string, proof []byte) (string, error) {
	s := struct {
		Nonce  string `json:"nonce,omitempty"`
		Sender string `json:"address,omitempty"`
	}{}
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		return "", err
	}
	n, ok := new(big.Int).SetString(strings.TrimPrefix(s.Nonce, "0x"), 16)
	if !ok {
		return "", errors.New("nonce format error")
	}
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return "", errors.Wrap(err, "Failed to connect to the Ethereum client")
	}
	privateKeyECDSA, err := crypto.HexToECDSA(r.operatorPrivateKeyECDSA)
	if err != nil {
		return "", errors.Wrap(err, "Error in parsing private key")
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "Failed to get chainID")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create transactor")
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := powerc20.NewPowerc20(contractAddr, client)
	if err != nil {
		return "", errors.Wrap(err, "Failed to instantiate a Token contract")
	}

	tx, err := contract.Mine(auth, n, common.HexToAddress(s.Sender), proof)
	if err != nil {
		return "", errors.Wrap(err, "Failed to submit mine transaction")
	}
	return tx.Hash().Hex(), nil
}

func (r *Processor) reportFail(taskID string, err error, topic *pubsub.Topic) {
	j, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    taskID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("json marshal p2p task state log data failed", "error", err, "taskID", taskID)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish task state log data to p2p network failed", "error", err, "taskID", taskID)
	}
}

func (r *Processor) reportSuccess(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {
	j, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    taskID,
			State:     state,
			Comment:   comment,
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("json marshal p2p task state log data failed", "error", err, "taskID", taskID)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish task state log data to p2p network failed", "error", err, "taskID", taskID)
	}
}

func (r *Processor) buildProjectOutputConfig(cfg project.Config) (outCfg output.Config) {
	var projOutCfg = cfg.Output

	switch projOutCfg.Type {
	case types.OutputEthereumContract:
		ethCfg := projOutCfg.Ethereum
		outCfg = output.NewEthereumContractConfig(chain.Name(ethCfg.ChainName), ethCfg.ContractAddress, r.operatorPrivateKeyECDSA)
	case types.OutputSolanaProgram:
		solCfg := projOutCfg.Solana
		outCfg = output.NewSolanaProgramConfig(chain.Name(solCfg.ChainName), solCfg.ProgramID, r.operatorPrivateKeyED25519, solCfg.StateAccountPK)
	default:
		outCfg = output.NewStdoutConfig()
	}
	return outCfg
}

func (r *Processor) Run() {
	// TODO project load & delete
}

func NewProcessor(vmHandler *vm.Handler, projectManager *project.Manager, outputFactory *output.Factory, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, iotexChainID int) (*Processor, error) {
	p := &Processor{
		vmHandler:                 vmHandler,
		operatorPrivateKeyECDSA:   operatorPrivateKey,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
		projectManager:            projectManager,
		outputFactory:             outputFactory,
	}

	ps, err := p2p.NewPubSubs(p.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	p.ps = ps

	for _, id := range projectManager.GetAllProjectID() {
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "add project %d pubsub failed", id)
		}
	}
	return p, nil
}
