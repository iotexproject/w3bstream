package vm

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/iotexproject/w3bstream/task"
)

var (
	// TODO: use real project IDs
	_pebbleProjectID = big.NewInt(923)
	_geoProjectID    = big.NewInt(942)
)

func LoadPayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
	switch task.ProjectID.String() {
	case _pebbleProjectID.String():
		return encodePebblePayload(task, projectConfig)
	case _geoProjectID.String():
		panic("not implemented")
	default:
		return task.Payload, nil
	}
}

type ProofofLivenessCircuit struct {
	PayloadHash []uints.U8
	Timestamp   frontend.Variable `gnark:",public"`
	PubBytes    []uints.U8        `gnark:",public"`
	SigBytes    []uints.U8
}

func (circuit *ProofofLivenessCircuit) Define(api frontend.API) error { return nil }

func encodePebblePayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
	sig := task.Signature
	pubbytes := task.DevicePubKey
	h1, _, data, err := recover(task, projectConfig)
	if err != nil {
		return nil, err
	}
	timestamp := data[0].(uint64)
	payloadHash := h1

	assignment := ProofofLivenessCircuit{
		PayloadHash: uints.NewU8Array(payloadHash[:]),
		Timestamp:   timestamp,
		SigBytes:    uints.NewU8Array(sig[:]),
		PubBytes:    uints.NewU8Array(pubbytes),
	}
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

// TODO duplicate logic
func recover(t *task.Task, cfg *project.Config) (h1, h2 common.Hash, data []any, err error) {
	req := api.CreateTaskReq{
		Nonce:          t.Nonce,
		ProjectID:      t.ProjectID.String(),
		ProjectVersion: t.ProjectVersion,
		Payload:        t.Payload,
	}
	reqJson, err := json.Marshal(req)
	if err != nil {
		return common.Hash{}, common.Hash{}, nil, errors.Wrap(err, "failed to marshal request into json format")
	}

	switch cfg.HashAlgorithm {
	default:
		h1 = sha256.Sum256(reqJson)
		if len(cfg.SignedKeys) == 0 {
			h2 = h1
		} else {
			// concatenate payload hash with signed keys from payload json
			if ok := gjson.ValidBytes(req.Payload); !ok {
				return common.Hash{}, common.Hash{}, nil, errors.Wrap(err, "failed to validate payload in json format")
			}
			buf := new(bytes.Buffer)
			buf.Write(h1[:])
			for _, k := range cfg.SignedKeys {
				value := gjson.GetBytes(req.Payload, k.Name)
				switch k.Type {
				case "uint64":
					data = append(data, value.Uint())
					if err := binary.Write(buf, binary.LittleEndian, value.Uint()); err != nil {
						return common.Hash{}, common.Hash{}, nil, errors.New("failed to convert uint64 to bytes array")
					}
				}
			}
			h2 = sha256.Sum256(buf.Bytes())
		}
	}
	return
}
