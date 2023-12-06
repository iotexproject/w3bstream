package adapter

import (
	"errors"
	"log/slog"

	solcommon "github.com/blocto/solana-go-sdk/common"
	soltypes "github.com/blocto/solana-go-sdk/types"

	"github.com/machinefi/sprout/output/chain/solana"
)

type (
	// SolanaProgram is the solana program adapter
	SolanaProgram struct {
		endpoint       string
		programID      string
		secretKey      string
		stateAccountPK string
	}

	// SolanaProgramResult is the result of the solana program adapter
	SolanaProgramResult struct {
		TxHash string
	}
)

// NewSolanaProgram returns a new solana program adapter
func NewSolanaProgram(endpoint, programID, secretKey, stateAccountPK string) (*SolanaProgram, error) {
	if len(secretKey) == 0 {
		return nil, errors.New("secretkey is empty")
	}
	return &SolanaProgram{
		endpoint:       endpoint,
		programID:      programID,
		secretKey:      secretKey,
		stateAccountPK: stateAccountPK,
	}, nil
}

// Output outputs the proof to the ethereum contract
func (e *SolanaProgram) Output(proof []byte) (Result, error) {
	slog.Debug("outputing to solana program", "chain endpoint", e.endpoint)
	// pack instructions
	ins := e.packInstructions(proof[:8])

	// send tx
	txHash, err := solana.SendTX(e.endpoint, e.secretKey, ins)
	if err != nil {
		return nil, err
	}
	slog.Debug("output success", "txHash", txHash)

	return &SolanaProgramResult{TxHash: txHash}, nil
}

// encodeData encodes the proof into the data field of the instruction
// the first byte is the instruction, which is 0 for now;
// the rest is the proof data.
// e.g. assume proof is [0x01, 0x02, 0x03], then the encoded data is [0x00, 0x01, 0x02, 0x03]
func (e *SolanaProgram) encodeData(proof []byte) []byte {
	data := []byte{}
	data = append(data, byte(0)) // 0 means submit proof
	data = append(data, proof...)
	return data
}

func (e *SolanaProgram) packInstructions(proof []byte) []soltypes.Instruction {
	accounts := []soltypes.AccountMeta{}
	if e.stateAccountPK != "" {
		accounts = append(accounts, soltypes.AccountMeta{
			PubKey:     solcommon.PublicKeyFromString(e.stateAccountPK),
			IsSigner:   false,
			IsWritable: true,
		})
	}

	return []soltypes.Instruction{
		{
			ProgramID: solcommon.PublicKeyFromString(e.programID),
			Accounts:  accounts,
			Data:      e.encodeData(proof),
		},
	}
}
