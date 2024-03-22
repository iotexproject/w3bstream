package persistence

import (
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/persistence/znode"
)

type Prover struct {
	mux       sync.Mutex
	proverIDs map[string]bool

	contractAddress string
	chainEndpoint   string
}

func (z *Prover) GetAll() []string {
	z.mux.Lock()
	defer z.mux.Unlock()

	ids := []string{}
	for d := range z.proverIDs {
		ids = append(ids, d)
	}
	return ids
}

// TODO monitor prover contract event
func NewProver(chainEndpoint, contractAddress string) (*Prover, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	}
	instance, err := znode.NewZnode(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "new prover contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	proverIDs := map[string]bool{}

	for i := uint64(1); ; i++ {
		prover, err := instance.Znodes(nil, i)
		if err != nil {
			slog.Error("get prover from chain failed", "prover_id", i, "error", err)
			continue
		}
		if prover.Did == "" {
			break
		}
		proverIDs[prover.Did] = true
	}

	return &Prover{
		proverIDs:       proverIDs,
		contractAddress: contractAddress,
		chainEndpoint:   chainEndpoint,
	}, nil
}
