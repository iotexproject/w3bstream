package task

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Task struct {
	ID             common.Hash    `json:"id"`
	ProjectID      uint64         `json:"projectID"`
	ProjectVersion string         `json:"projectVersion"`
	DeviceID       common.Address `json:"deviceID"`
	Payloads       [][]byte       `json:"payloads"`
	Signature      []byte         `json:"signature,omitempty"`
}

func (t *Task) Sign(prv *ecdsa.PrivateKey) ([]byte, error) {
	h, err := t.hash()
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(h.Bytes(), prv)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign")
	}
	return sig, nil
}

func (t *Task) VerifySignature(pubKey []byte) error {
	h, err := t.hash()
	if err != nil {
		return err
	}
	sigpk, err := crypto.Ecrecover(h.Bytes(), t.Signature)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key")
	}
	if !bytes.Equal(sigpk, pubKey) {
		return errors.New("task signature unmatched")
	}
	return nil
}

func (t *Task) hash() (common.Hash, error) {
	nt := *t
	nt.Signature = nil
	j, err := json.Marshal(&nt)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to marshal task")
	}

	return crypto.Keccak256Hash(j), nil
}

type StateLog struct {
	TaskID    uint64
	ProjectID uint64
	State     State
	Comment   string
	Result    []byte
	Signature string
	ProverID  uint64
	CreatedAt time.Time
}

func (l *StateLog) SignerAddress(task *Task) (common.Address, error) {
	buf := bytes.NewBuffer(nil)

	if err := binary.Write(buf, binary.BigEndian, task.ID); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write binary")
	}
	if err := binary.Write(buf, binary.BigEndian, task.ProjectID); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write binary")
	}
	if _, err := buf.Write(task.DeviceID[:]); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}
	if _, err := buf.Write(crypto.Keccak256Hash(task.Payloads...).Bytes()); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}
	if _, err := buf.Write(l.Result); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), task.Signature)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to recover public key")
	}

	publicKey, err := crypto.UnmarshalPubkey(sigpk)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to unmarshal public key")
	}
	return crypto.PubkeyToAddress(*publicKey), nil
}

type State uint8

const (
	StateInvalid State = iota
	StatePacked
	StateDispatched
	_
	StateProved
	_
	StateOutputted
	StateFailed
)

func (s State) String() string {
	switch s {
	case StatePacked:
		return "packed"
	case StateDispatched:
		return "dispatched"
	case StateProved:
		return "proved"
	case StateOutputted:
		return "outputted"
	case StateFailed:
		return "failed"
	default:
		return "invalid"
	}
}
