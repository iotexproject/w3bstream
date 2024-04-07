package types

import (
	"bytes"
	"encoding/binary"
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
	ClientID       string   `json:"clientID"`
	Signature      string   `json:"signature"`
}

func (t *Task) VerifySignature(pubkey []byte) error {
	sig, err := hexutil.Decode(t.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to decode task signature")
	}

	buf := bytes.NewBuffer(nil)

	if err = binary.Write(buf, binary.BigEndian, t.ID); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, t.ProjectID); err != nil {
		return err
	}
	if _, err = buf.WriteString(t.ClientID); err != nil {
		return err
	}
	if _, err = buf.Write(crypto.Keccak256Hash(t.Data...).Bytes()); err != nil {
		return err
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, pubkey) {
		return errors.New("task signature unmatched")
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
	ProjectID uint64
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
		return errors.Wrap(err, "failed to decode task signature")
	}

	buf := bytes.NewBuffer(nil)

	if err = binary.Write(buf, binary.BigEndian, task.ID); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, task.ProjectID); err != nil {
		return err
	}
	if _, err = buf.WriteString(task.ClientID); err != nil {
		return err
	}
	if _, err = buf.Write(crypto.Keccak256Hash(task.Data...).Bytes()); err != nil {
		return err
	}
	if _, err = buf.Write(l.Result); err != nil {
		return err
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, proverPubKey) {
		return errors.New("proof signature unmatched")
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
