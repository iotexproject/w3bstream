package message

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/output/chain"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
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
	return p, nil
}

func (r *Processor) Run() {
	go r.runMessageRequest()
}

// TODO support batch message fetch & fetch frequency define
func (r *Processor) runMessageRequest() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	projectIDs := r.projectManager.GetAllProjectID()

	for range ticker.C {
		for _, projectID := range projectIDs {
			if err := r.ps.Add(projectID); err != nil {
				slog.Error("add project pubsub failed", "error", err)
				continue
			}

			d := &p2p.Data{
				Request: &p2p.RequestData{
					ProjectID: projectID,
				},
			}
			if err := r.ps.Publish(projectID, d); err != nil {
				slog.Error("publish data failed", "error", err)
			}
		}
	}
}

func (r *Processor) handleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Message == nil || len(d.Message.Messages) == 0 {
		return
	}
	ms := d.Message.Messages
	mids := r.getMessageIDs(ms)
	slog.Debug("get new messages", "message_ids", mids)
	r.reportSuccess(mids, types.MessageStateFetched, "", topic)

	project, err := r.projectManager.Get(ms[0].ProjectID)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail(mids, err, topic)
		return
	}

	r.reportSuccess(mids, types.MessageStateProving, "", topic)
	res, err := r.vmHandler.Handle(ms, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail(mids, err, topic)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess(mids, types.MessageStateProved, string(res), topic)

	// output proof
	outCfg := r.buildProjectOutputConfig(project.Config)
	outputter, err := r.outputFactory.NewOutputter(outCfg)
	if err != nil {
		err = errors.Wrap(err, "fail to init outputter")
		slog.Error(err.Error())
		r.reportFail(mids, err, topic)
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", outputter))

	r.reportSuccess(mids, types.MessageStateOutputting, "output proof", topic)
	outRes, err := outputter.Output(res)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(mids, err, topic)
		return
	}
	r.reportSuccess(mids, types.MessageStateOutputted, fmt.Sprintf("output result: %+v", outRes), topic)
	slog.Debug("output success", "result", outRes)
}

func (r *Processor) reportFail(messageIDs []string, err error, topic *pubsub.Topic) {
	d := p2p.Data{
		Response: &p2p.ResponseData{
			MessageIDs: messageIDs,
			State:      types.MessageStateFailed,
			Comment:    err.Error(),
			CreatedAt:  time.Now(),
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) reportSuccess(messageIDs []string, state types.MessageState, comment string, topic *pubsub.Topic) {
	d := p2p.Data{
		Response: &p2p.ResponseData{
			MessageIDs: messageIDs,
			State:      state,
			Comment:    comment,
			CreatedAt:  time.Now(),
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) getMessageIDs(ms []*types.Message) []string {
	ids := []string{}
	for _, m := range ms {
		ids = append(ids, m.ID)
	}
	return ids
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
