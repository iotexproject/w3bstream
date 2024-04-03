package task

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/utils/distance"
	"github.com/machinefi/sprout/vm"
)

type VMHandler interface {
	Handle(task *types.Task, vmtype vm.Type, code string, expParam string) ([]byte, error)
}

type Processor struct {
	vmHandler        VMHandler
	projectManager   ProjectManager
	proverPrivateKey *ecdsa.PrivateKey
	sequencerPubKey  []byte
	proverID         string
	projectProvers   sync.Map
}

func (r *Processor) HandleProjectProvers(projectID uint64, provers []string) {
	r.projectProvers.Store(projectID, provers)
}

func (r *Processor) HandleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Task == nil {
		return
	}
	t := d.Task

	p, err := r.projectManager.Get(t.ProjectID)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", t.ProjectID)
		r.reportFail(t, err, topic)
		return
	}
	c, err := p.GetDefaultConfig()
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.ProjectID, "project_version", p.DefaultVersion)
		r.reportFail(t, err, topic)
		return
	}

	var provers []string
	proversValue, ok := r.projectProvers.Load(t.ProjectID)
	if ok {
		provers = proversValue.([]string)
	}
	if len(provers) > 1 {
		workProver := distance.GetMinNLocation(provers, t.ID, 1)
		if workProver[0] != r.proverID {
			slog.Info("the task not scheduld to this prover", "project_id", t.ProjectID, "task_id", t.ID)
			return
		}
	}

	if err := t.VerifySignature(r.sequencerPubKey); err != nil {
		slog.Error("failed to verify task sign", "error", err)
		return
	}

	slog.Debug("get a new task", "task_id", t.ID)
	r.reportSuccess(t, types.TaskStateDispatched, nil, "", topic)

	res, err := r.vmHandler.Handle(t, c.VMType, c.Code, c.CodeExpParam)
	if err != nil {
		slog.Error("failed to generate proof", "error", err)
		r.reportFail(t, err, topic)
		return
	}
	signature, err := r.signProof(t, res)
	if err != nil {
		slog.Error("failed to sign proof", "error", err)
		r.reportFail(t, err, topic)
		return
	}

	r.reportSuccess(t, types.TaskStateProved, res, signature, topic)
}

func (r *Processor) signProof(t *types.Task, res []byte) (string, error) {
	buf := bytes.NewBuffer([]byte(fmt.Sprintf("%d%d%s", t.ID, t.ProjectID, t.ClientID)))
	for _, v := range t.Data {
		buf.Write(v)
	}
	buf.Write(res)

	h := crypto.Keccak256Hash(buf.Bytes())
	sig, err := crypto.Sign(h.Bytes(), r.proverPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign proof")
	}
	return hexutil.Encode(sig), nil
}

func (r *Processor) reportFail(t *types.Task, err error, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    t.ID,
			ProjectID: t.ProjectID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("failed to marshal p2p task state log data to json", "error", err, "task_id", t.ID)
		return
	}
	if err := topic.Publish(context.Background(), d); err != nil {
		slog.Error("failed to publish task state log data to p2p network", "error", err, "task_id", t.ID)
	}
}

func (r *Processor) reportSuccess(t *types.Task, state types.TaskState, result []byte, signature string, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    t.ID,
			ProjectID: t.ProjectID,
			State:     state,
			Result:    result,
			Signature: signature,
			ProverID:  r.proverID,
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("failed to marshal p2p task state log data to json", "error", err, "task_id", t.ID)
		return
	}
	if err := topic.Publish(context.Background(), d); err != nil {
		slog.Error("failed to publish task state log data to p2p network", "error", err, "task_id", t.ID)
	}
}

func NewProcessor(vmHandler VMHandler, projectManager ProjectManager, proverPrivateKey *ecdsa.PrivateKey, seqPubkey []byte, proverID string) *Processor {
	return &Processor{
		vmHandler:        vmHandler,
		projectManager:   projectManager,
		proverPrivateKey: proverPrivateKey,
		sequencerPubKey:  seqPubkey,
		proverID:         proverID,
	}
}
