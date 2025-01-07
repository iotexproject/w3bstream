package vm

import (
	_ "embed"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/iotexproject/w3bstream/task"
)

var (
	_pebbleProjectID = big.NewInt(923)
	_geoProjectID    = big.NewInt(942)
)

func LoadPayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
	// switch task.ProjectID.String() {
	// case _pebbleProjectID.String():
	return encodePebblePayload(task, projectConfig)
	// case _geoProjectID.String():
	// 	return encodeGeodnetPayload(task, projectConfig)
	// default:
	// 	return task.Payload, nil
	// }
}

type ProofofLivenessCircuit struct {
	PayloadHash []uints.U8
	Timestamp   frontend.Variable `gnark:",public"`
	PubBytes    []uints.U8        `gnark:",public"`
	SigBytes    []uints.U8
}

func (circuit *ProofofLivenessCircuit) Define(api frontend.API) error { return nil }

func encodePebblePayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
	sig := task.Signature[:64]
	pubbytes := task.DevicePubKey
	payloadHash, _, _, data, err := api.HashTask(
		&api.CreateTaskReq{
			Nonce:          task.Nonce,
			ProjectID:      task.ProjectID.String(),
			ProjectVersion: task.ProjectVersion,
			Payload:        task.Payload,
		}, projectConfig)
	if err != nil {
		return nil, err
	}
	timestamp := data[0].(uint64)

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

type ProofofMovenessCircuit struct {
	LastPayloadHash []uints.U8
	LastTimestamp   frontend.Variable
	LastLatitude    frontend.Variable
	LastLongitude   frontend.Variable
	LastSigBytes    []uints.U8

	CurPayloadHash []uints.U8
	CurTimestamp   frontend.Variable `gnark:",public"`
	CurLatitude    frontend.Variable
	CurLongitude   frontend.Variable
	CurSigBytes    []uints.U8

	IsMoved frontend.Variable `gnark:",public"`

	PubBytes []uints.U8 `gnark:",public"`
}

func (circuit *ProofofMovenessCircuit) Define(api frontend.API) error { return nil }

func encodeGeodnetPayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
	if task.PrevTask == nil {
		return nil, errors.New("geodnet project miss previous task")
	}
	lastPayloadHash, _, _, lastData, err := api.HashTask(
		&api.CreateTaskReq{
			Nonce:          task.PrevTask.Nonce,
			ProjectID:      task.PrevTask.ProjectID.String(),
			ProjectVersion: task.PrevTask.ProjectVersion,
			Payload:        task.PrevTask.Payload,
		}, projectConfig)
	if err != nil {
		return nil, err
	}
	curPayloadHash, _, _, curData, err := api.HashTask(
		&api.CreateTaskReq{
			Nonce:          task.Nonce,
			ProjectID:      task.ProjectID.String(),
			ProjectVersion: task.ProjectVersion,
			Payload:        task.Payload,
		}, projectConfig)
	if err != nil {
		return nil, err
	}
	lastTimestamp := lastData[0].(uint64)
	lastLatitude := lastData[1].(uint64)
	lastLongitude := lastData[2].(uint64)
	lastSig := task.PrevTask.Signature[:64]
	curTimestamp := curData[0].(uint64)
	curLatitude := curData[1].(uint64)
	curLongitude := curData[2].(uint64)
	curSig := task.Signature[:64]
	isMove := uint64(0)
	if (abs(lastLatitude, curLatitude) > 10^3) || (abs(lastLongitude, curLongitude) > 10^3) {
		isMove = 1
	}

	assignment := ProofofMovenessCircuit{
		LastPayloadHash: uints.NewU8Array(lastPayloadHash[:]),
		LastTimestamp:   lastTimestamp,
		LastLatitude:    lastLatitude,
		LastLongitude:   lastLongitude,
		LastSigBytes:    uints.NewU8Array(lastSig[:]),

		CurPayloadHash: uints.NewU8Array(curPayloadHash[:]),
		CurTimestamp:   curTimestamp,
		CurLatitude:    curLatitude,
		CurLongitude:   curLongitude,
		CurSigBytes:    uints.NewU8Array(curSig[:]),

		IsMoved:  isMove,
		PubBytes: uints.NewU8Array(task.DevicePubKey),
	}
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func abs(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}
