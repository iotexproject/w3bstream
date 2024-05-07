package task

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	sig, err := hexutil.Decode(task.Signature)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to decode task signature")
	}

	buf := bytes.NewBuffer(nil)

	if err = binary.Write(buf, binary.BigEndian, task.ID); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write binary")
	}
	if err = binary.Write(buf, binary.BigEndian, task.ProjectID); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write binary")
	}
	if _, err = buf.WriteString(task.ClientID); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}
	if _, err = buf.Write(crypto.Keccak256Hash(task.Data...).Bytes()); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}
	if _, err = buf.Write(l.Result); err != nil {
		return common.Address{}, errors.Wrap(err, "failed to write bytes buffer")
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sigpk, err := crypto.Ecrecover(h.Bytes(), sig)
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
