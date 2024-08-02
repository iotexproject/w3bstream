package processor

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"log/slog"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/util/distance"
)

type VMHandler interface {
	Handle(task *task.Task, vmTypeID uint64, code string, expParams []string) ([]byte, error)
}

type Project func(projectID uint64) (*project.Project, error)

type Processor struct {
	vmHandler               VMHandler
	project                 Project
	proverPrivateKey        *ecdsa.PrivateKey
	defaultDatasourcePubKey []byte
	proverID                uint64
	projectProvers          sync.Map
}

func (r *Processor) HandleProjectProvers(projectID uint64, proverIDs []uint64) {
	r.projectProvers.Store(projectID, proverIDs)
}

func (r *Processor) HandleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Task == nil {
		return
	}
	t := d.Task

	p, err := r.project(t.ProjectID)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", t.ProjectID)
		r.reportFail(t, err, topic)
		return
	}
	c, err := p.DefaultConfig()
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.ProjectID, "project_version", p.DefaultVersion)
		r.reportFail(t, err, topic)
		return
	}

	var provers []uint64
	proversValue, ok := r.projectProvers.Load(t.ProjectID)
	if ok {
		provers = proversValue.([]uint64)
	}
	if len(provers) > 1 {
		workProvers := distance.Sort(provers, t.ID)
		if workProvers[0] != r.proverID {
			slog.Info("the task not scheduled to this prover", "project_id", t.ProjectID, "task_id", t.ID)
			return
		}
	}

	pubKey := r.defaultDatasourcePubKey
	if p.DatasourcePubKey != "" {
		pubKey, err = hexutil.Decode(p.DatasourcePubKey)
		if err != nil {
			slog.Error("failed to decode datasource public key", "error", err, "project_id", t.ProjectID)
			r.reportFail(t, err, topic)
			return
		}
	}

	if err := t.VerifySignature(pubKey); err != nil {
		slog.Error("failed to verify task signature", "error", err)
		return
	}

	slog.Debug("get a new task", "project_id", t.ProjectID, "task_id", t.ID)
	r.reportSuccess(t, task.StateDispatched, nil, "", topic)

	res, err := r.vmHandler.Handle(t, c.VMTypeID, c.Code, c.CodeExpParams)
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

	r.reportSuccess(t, task.StateProved, res, signature, topic)
}

func (r *Processor) signProof(t *task.Task, res []byte) (string, error) {
	// TODO: use protobuf or json to encode
	buf := bytes.NewBuffer(nil)

	if err := binary.Write(buf, binary.BigEndian, t.ID); err != nil {
		return "", err
	}
	if err := binary.Write(buf, binary.BigEndian, t.ProjectID); err != nil {
		return "", err
	}
	if _, err := buf.WriteString(t.ClientID); err != nil {
		return "", err
	}
	if _, err := buf.Write(crypto.Keccak256Hash(t.Data...).Bytes()); err != nil {
		return "", err
	}
	if _, err := buf.Write(res); err != nil {
		return "", err
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sig, err := crypto.Sign(h.Bytes(), r.proverPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign proof")
	}
	return hexutil.Encode(sig), nil
}

func (r *Processor) reportFail(t *task.Task, err error, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2p.Data{
		TaskStateLog: &task.StateLog{
			TaskID:    t.ID,
			ProjectID: t.ProjectID,
			State:     task.StateFailed,
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

func (r *Processor) reportSuccess(t *task.Task, state task.State, result []byte, signature string, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2p.Data{
		TaskStateLog: &task.StateLog{
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

func NewProcessor(vmHandler VMHandler, project Project, proverPrivateKey *ecdsa.PrivateKey, defaultDatasourcePubKey []byte, proverID uint64) *Processor {
	return &Processor{
		vmHandler:               vmHandler,
		project:                 project,
		proverPrivateKey:        proverPrivateKey,
		defaultDatasourcePubKey: defaultDatasourcePubKey,
		proverID:                proverID,
	}
}
