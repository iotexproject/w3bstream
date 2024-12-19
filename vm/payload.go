package vm

import (
	_ "embed"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task"
)

var (
	// TODO: use real project IDs
	_pebbleProjectID = big.NewInt(1234)
	_geoProjectID    = big.NewInt(5678)
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
	// TODO: load from arguments
	payloadHash := []byte{}
	timestamp := 0
	sig := []byte{}
	pubbytes := []byte{}

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
