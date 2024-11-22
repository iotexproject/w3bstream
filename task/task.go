package task

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Task struct {
	ID             common.Hash    `json:"id"`
	ProjectID      uint64         `json:"projectID"`
	ProjectVersion string         `json:"projectVersion,omitempty"`
	DeviceID       common.Address `json:"deviceID"`
	Payloads       [][]byte       `json:"payloads"`
	Signature      []byte         `json:"signature,omitempty"`
}

func (t *Task) Sign(prv *ecdsa.PrivateKey) ([]byte, error) {
	h, err := t.Hash()
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(h.Bytes(), prv)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign")
	}
	sig[64] += 27
	return sig, nil
}

func (t *Task) VerifySignature(pubKey []byte) error {
	h, err := t.Hash()
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

func (t *Task) Hash() (common.Hash, error) {
	nt := *t
	nt.Signature = nil
	j, err := json.Marshal(&nt)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to marshal task")
	}

	return crypto.Keccak256Hash(j), nil
}
