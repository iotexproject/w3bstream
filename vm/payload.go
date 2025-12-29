package vm

import (
	_ "embed"
	"log/slog"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/iotexproject/w3bstream/task"
)

func loadPayload(tasks []*task.Task, projectConfig *project.Config) ([]byte, error) {
	switch projectConfig.ProofType {
	case "liveness":
		return encodeLivenessPayload(tasks[0], projectConfig)
	case "movement":
		return encodeMovementPayload(tasks, projectConfig)
	case "sum":
		return encodeSumPayload(tasks, projectConfig)
	default:
		return tasks[0].Payload, nil
	}
}

type ProofofLivenessCircuit struct {
	PayloadHash []uints.U8
	Timestamp   frontend.Variable `gnark:",public"`
	PubBytes    []uints.U8        `gnark:",public"`
	SigBytes    []uints.U8
}

func (circuit *ProofofLivenessCircuit) Define(api frontend.API) error { return nil }

func encodeLivenessPayload(task *task.Task, projectConfig *project.Config) ([]byte, error) {
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

type ProofOfMovementBatchCircuit struct {
	LastPayloadHash [10][32]uints.U8
	LastTimestamp   [10]frontend.Variable
	LastLatitude    [10]frontend.Variable
	LastLongitude   [10]frontend.Variable
	LastSigBytes    [10][64]uints.U8

	CurPayloadHash [10][32]uints.U8
	CurTimestamp   [10]frontend.Variable
	CurLatitude    [10]frontend.Variable
	CurLongitude   [10]frontend.Variable
	CurSigBytes    [10][64]uints.U8

	PubBytes [10][65]uints.U8

	EthAddress [10]frontend.Variable `gnark:",public"`
	IsMoved    frontend.Variable     `gnark:",public"`
}

func (circuit *ProofOfMovementBatchCircuit) Define(api frontend.API) error { return nil }

func encodeMovementPayload(tasks []*task.Task, projectConfig *project.Config) ([]byte, error) {
	if len(tasks) != 10 {
		return nil, errors.Errorf("invalid tasks len, expect %d, get %d", 10, len(tasks))
	}
	assignment := ProofOfMovementBatchCircuit{}
	movedFlags := []bool{}
	for i := range tasks {
		task := tasks[i]
		if task.PrevTask == nil {
			return nil, errors.New("movement project miss previous task")
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
		if (abs(lastLatitude, curLatitude) > 1_000) || (abs(lastLongitude, curLongitude) > 1_000) {
			isMove = 1
		}
		movedFlags = append(movedFlags, isMove > 0)

		assignment.LastPayloadHash[i] = [32]uints.U8(uints.NewU8Array(lastPayloadHash[:]))
		assignment.LastTimestamp[i] = lastTimestamp
		assignment.LastLatitude[i] = lastLatitude
		assignment.LastLongitude[i] = lastLongitude
		assignment.LastSigBytes[i] = [64]uints.U8(uints.NewU8Array(lastSig[:]))
		assignment.CurPayloadHash[i] = [32]uints.U8(uints.NewU8Array(curPayloadHash[:]))
		assignment.CurTimestamp[i] = curTimestamp
		assignment.CurLatitude[i] = curLatitude
		assignment.CurLongitude[i] = curLongitude
		assignment.CurSigBytes[i] = [64]uints.U8(uints.NewU8Array(curSig[:]))
		assignment.PubBytes[i] = [65]uints.U8(uints.NewU8Array(task.DevicePubKey))

		pubkey, err := crypto.UnmarshalPubkey(task.DevicePubKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal pubkey")
		}
		assignment.EthAddress[i] = crypto.PubkeyToAddress(*pubkey).Big()
	}

	var isMovedValue uint64 = 0
	for i := 0; i < 10; i++ {
		if movedFlags[i] {
			isMovedValue |= (1 << i)
		}
	}
	assignment.IsMoved = isMovedValue

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

const SumMaxItems = 2

type ProofOfSumCircuit struct {
	PayloadHashs [SumMaxItems][32]uints.U8
	Timestamps   [SumMaxItems]frontend.Variable
	Values       [SumMaxItems]frontend.Variable
	SigBytes     [SumMaxItems][64]uints.U8

	PubBytes  [SumMaxItems][65]uints.U8
	StartTime frontend.Variable

	Threshold  frontend.Variable `gnark:",public"`
	EthAddress frontend.Variable `gnark:",public"`
}

func (circuit *ProofOfSumCircuit) Define(api frontend.API) error { return nil }

func encodeSumPayload(tasks []*task.Task, projectConfig *project.Config) ([]byte, error) {
	if len(tasks) != 1 {
		return nil, errors.Errorf("invalid tasks len, expect %d, get %d", 1, len(tasks))
	}
	assignment := ProofOfSumCircuit{}
	task := tasks[0]
	if task.PrevTask == nil {
		return nil, errors.New("sum project miss previous task")
	}

	slog.Debug("--------------1")
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
	slog.Debug("--------------2")
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
	slog.Debug("--------------3")
	lastTimestamp := lastData[0].(uint64)
	lastValue := lastData[1].(uint64)
	lastSig := task.PrevTask.Signature[:64]
	curTimestamp := curData[0].(uint64)
	curValue := curData[1].(uint64)
	curSig := task.Signature[:64]

	slog.Debug("sum payload", "lastTimestamp", lastTimestamp, "lastValue", lastValue, "lastSig", lastSig,
		"curTimestamp", curTimestamp, "curValue", curValue, "curSig", curSig, "lastPayloadHash", lastPayloadHash,
		"curPayloadHash", curPayloadHash)

	assignment.PayloadHashs[0] = [32]uints.U8(uints.NewU8Array(lastPayloadHash[:]))
	assignment.Timestamps[0] = lastTimestamp
	assignment.Values[0] = lastValue
	assignment.SigBytes[0] = [64]uints.U8(uints.NewU8Array(lastSig[:]))
	assignment.PayloadHashs[1] = [32]uints.U8(uints.NewU8Array(curPayloadHash[:]))
	assignment.Timestamps[1] = curTimestamp
	assignment.Values[1] = curValue
	assignment.SigBytes[1] = [64]uints.U8(uints.NewU8Array(curSig[:]))
	assignment.PubBytes[0] = [65]uints.U8(uints.NewU8Array(task.DevicePubKey))
	assignment.PubBytes[1] = [65]uints.U8(uints.NewU8Array(task.DevicePubKey))
	assignment.Threshold = uint64(10)

	slog.Debug("--------------5")

	pubkey, err := crypto.UnmarshalPubkey(task.DevicePubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal pubkey")
	}
	assignment.EthAddress = crypto.PubkeyToAddress(*pubkey).Big()

	slog.Debug("--------------6")

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, errors.Wrap(err, "failed to new witness")
	}

	slog.Debug("--------------7")

	data, err := witness.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal witness binary")
	}
	slog.Debug("--------------8")

	return data, nil
}
