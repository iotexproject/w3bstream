package task

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
	Sign           string   `json:"sign"`
}

func (t *Task) verify(pubkey []byte) error {
	sig, err := hexutil.Decode(t.Sign)
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
	Task       Task
	State      TaskState
	Comment    string
	Result     []byte
	SignResult string
	// TODO del
	proverID  string
	CreatedAt time.Time
}

func (l *TaskStateLog) verify(pubkey []byte) error {
	sig, err := hexutil.Decode(l.Task.Sign)
	if err != nil {
		return errors.Wrap(err, "failed to decode task sign")
	}

	data := bytes.NewBuffer([]byte(fmt.Sprintf("%d%d%s", l.Task.ID, l.Task.ProjectID, l.Task.ClientDID)))
	for _, v := range l.Task.Data {
		data.Write(v)
	}
	data.Write(l.Result)

	h := crypto.Keccak256Hash(data.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, pubkey) {
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

type p2pData struct {
	Task         *Task         `json:"task,omitempty"`
	TaskStateLog *TaskStateLog `json:"taskStateLog,omitempty"`
}
