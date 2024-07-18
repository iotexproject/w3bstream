package contract

import (
	"math/big"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/smartcontracts/go/prover"
)

func TestNewProver(t *testing.T) {
	r := require.New(t)
	p := newProject()
	r.Equal(p.Paused, true)
}

func TestProver_Merge(t *testing.T) {
	r := require.New(t)

	np := &Prover{}
	paused := true
	addr := common.Address{}
	nodeType := uint64(10)
	diff := &proverDiff{
		id:              1,
		operatorAddress: &addr,
		paused:          &paused,
		nodeTypes:       &nodeType,
	}
	np.merge(diff)
	r.Equal(np.ID, diff.id)
	r.Equal(np.NodeTypes, *diff.nodeTypes)
}

func TestBlockProver_Merge(t *testing.T) {
	r := require.New(t)

	np := &blockProver{Provers: map[uint64]*Prover{}}

	paused := true
	addr := common.Address{}
	nodeType := uint64(10)
	diff := &blockProverDiff{
		diffs: map[uint64]*proverDiff{
			1: {
				id:              1,
				operatorAddress: &addr,
				paused:          &paused,
				nodeTypes:       &nodeType,
			},
		},
	}

	np.merge(diff)
	r.Equal(len(np.Provers), 1)
}

func TestContract_processProverLogs(t *testing.T) {
	r := require.New(t)
	id := new(big.Int).SetUint64(1)
	filterer := &prover.ProverFilterer{}
	c := &Contract{
		proverInstance: &prover.Prover{ProverFilterer: *filterer},
	}

	t.Run("FailedToParseProverOperatorSetEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseOperatorSet", &prover.ProverOperatorSet{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{operatorSetTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		_, err := c.processProverLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverNodeTypeUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseNodeTypeUpdated", &prover.ProverNodeTypeUpdated{Id: id, Typ: id}, errors.New(t.Name()))
		logs := []types.Log{
			{
				Topics:      []common.Hash{nodeTypeUpdatedTopic},
				BlockNumber: 99,
				TxIndex:     1,
			},
		}

		_, err := c.processProverLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverPausedUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProverPaused", &prover.ProverProverPaused{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{proverPausedTopic},
				BlockNumber: 100,
				TxIndex:     2,
			},
		}

		_, err := c.processProverLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverResumedUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProverResumed", &prover.ProverProverResumed{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{proverResumedTopic},
				BlockNumber: 101,
				TxIndex:     1,
			},
		}

		_, err := c.processProverLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseOperatorSet", &prover.ProverOperatorSet{Id: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseNodeTypeUpdated", &prover.ProverNodeTypeUpdated{Id: id, Typ: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProverPaused", &prover.ProverProverPaused{Id: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProverResumed", &prover.ProverProverResumed{Id: id}, nil)

		logs := []types.Log{
			{
				Topics:      []common.Hash{operatorSetTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{nodeTypeUpdatedTopic},
				BlockNumber: 99,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{proverPausedTopic},
				BlockNumber: 100,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{proverResumedTopic},
				BlockNumber: 101,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{proverResumedTopic},
				BlockNumber: 101,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{proverResumedTopic},
				BlockNumber: 98,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{proverResumedTopic},
				BlockNumber: 98,
				TxIndex:     1,
			},
		}

		diffs, err := c.processProverLogs(logs)
		r.NoError(err)
		r.Equal(len(diffs), 4)
	})
}
