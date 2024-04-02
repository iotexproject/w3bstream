package types

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Task struct {
	ID             uint64   `json:"id"`
	ProjectID      uint64   `json:"projectID"`
	ProjectVersion string   `json:"projectVersion"`
	Data           [][]byte `json:"data"`
	ClientDID      string   `json:"clientDID"`
	Signature      string   `json:"sign"`
}

func (t *Task) VerifySignature(pubkey []byte) error {
	sig, err := hexutil.Decode(t.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to decode task sign")
	}

	data := bytes.NewBuffer([]byte(fmt.Sprintf("%d%d%s", t.ID, t.ProjectID, t.ClientDID)))
	for _, v := range t.Data {
		data.Write(v)
	}
	h := crypto.Keccak256Hash(data.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, pubkey) {
		return errors.New("task sign unmatched")
	}
	return nil
}

type TaskState uint8

const (
	TaskStateInvalid TaskState = iota
	TaskStatePacked
	TaskStateDispatched
	_
	TaskStateProved
	_
	TaskStateOutputted
	TaskStateFailed
)

type TaskStateLog struct {
	TaskID    uint64
	State     TaskState
	Comment   string
	Result    []byte
	Signature string
	ProverID  string
	CreatedAt time.Time
}

func (l *TaskStateLog) VerifySignature(pubkey string, task *Task) error {
	proverPubKey, err := hexutil.Decode(pubkey)
	if err != nil {
		return errors.Wrap(err, "failed to decode prover pubkey")
	}

	sig, err := hexutil.Decode(task.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to decode task sign")
	}

	data := bytes.NewBuffer([]byte(fmt.Sprintf("%d%d%s", task.ID, task.ProjectID, task.ClientDID)))
	for _, v := range task.Data {
		data.Write(v)
	}
	data.Write(l.Result)

	h := crypto.Keccak256Hash(data.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, proverPubKey) {
		return errors.New("proof sign unmatched")
	}
	return nil
}

func (s TaskState) String() string {
	switch s {
	case TaskStatePacked:
		return "packed"
	case TaskStateDispatched:
		return "dispatched"
	case TaskStateProved:
		return "proved"
	case TaskStateOutputted:
		return "outputted"
	case TaskStateFailed:
		return "failed"
	default:
		return "invalid"
	}
}
