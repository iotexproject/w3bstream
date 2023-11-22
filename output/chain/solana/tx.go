package solana

import (
	"context"
	"crypto/ed25519"

	"github.com/blocto/solana-go-sdk/client"
	solcommon "github.com/blocto/solana-go-sdk/common"
	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func SendTX(endpoint, privateKey string, ins []soltypes.Instruction) (string, error) {
	cli := client.NewClient(endpoint)
	b := common.FromHex(privateKey)
	pk := ed25519.PrivateKey(b)
	account := soltypes.Account{
		PublicKey:  solcommon.PublicKeyFromBytes(pk.Public().(ed25519.PublicKey)),
		PrivateKey: pk,
	}
	if len(ins) == 0 {
		return "", errors.New("missing instruction data")
	}

	resp, err := cli.GetLatestBlockhash(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "failed to get solana latest block hash")
	}
	tx, err := soltypes.NewTransaction(soltypes.NewTransactionParam{
		Message: soltypes.NewMessage(soltypes.NewMessageParam{
			FeePayer:        account.PublicKey,
			RecentBlockhash: resp.Blockhash,
			Instructions:    ins,
		}),
		Signers: []soltypes.Account{account},
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to build solana raw tx")
	}

	hash, err := cli.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", errors.Wrap(err, "failed to send solana tx")
	}
	return hash, nil
}
