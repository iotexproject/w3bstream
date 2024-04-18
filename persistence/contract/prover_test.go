package contract

import (
	"math/big"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/smartcontracts/go/prover"
	"github.com/machinefi/sprout/util/hash"
)

func TestProver_Merge(t *testing.T) {
	r := require.New(t)

	np := &Prover{}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &Prover{
		ID:              1,
		BlockNumber:     100,
		Paused:          &paused,
		OperatorAddress: common.Address{1},
		NodeTypes:       1,
	}
	np.Merge(diff)
	r.Equal(np, diff)
}

func TestBlockProver_Merge(t *testing.T) {
	r := require.New(t)

	np := &BlockProver{Provers: map[uint64]*Prover{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &BlockProver{
		BlockNumber: 100,
		Provers: map[uint64]*Prover{
			1: {
				ID:              1,
				BlockNumber:     100,
				Paused:          &paused,
				OperatorAddress: common.Address{1},
				NodeTypes:       1,
			},
		},
	}

	np.Merge(diff)
	r.Equal(np, diff)
}

func TestListAndWatchProver(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToDialChain", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := ListAndWatchProver("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProverContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, errors.New(t.Name()))

		_, err := ListAndWatchProver("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToQueryChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(prover.NewProver, &prover.Prover{}, nil)
		p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1), errors.New(t.Name()))

		_, err := ListAndWatchProver("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(prover.NewProver, &prover.Prover{}, nil)
		p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1), nil)
		p.ApplyFuncReturn(listProver, nil)
		p.ApplyFuncReturn(watchProver)

		_, err := ListAndWatchProver("", "", 0)
		r.NoError(err)
	})
}

func TestListProver(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToGetOperator", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		caller := &prover.ProverCaller{}
		p.ApplyMethodReturn(caller, "Operator", nil, errors.New(t.Name()))

		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToGetPaused", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		caller := &prover.ProverCaller{}
		p.ApplyMethodReturn(caller, "Operator", common.Address{}, nil)
		p.ApplyMethodReturn(caller, "IsPaused", false, errors.New(t.Name()))

		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToGetNodeType", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		caller := &prover.ProverCaller{}
		p.ApplyMethodReturn(caller, "Operator", common.Address{}, nil)
		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
		p.ApplyMethodReturn(caller, "NodeType", nil, errors.New(t.Name()))

		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		caller := &prover.ProverCaller{}
		p.ApplyMethodSeq(caller, "Operator", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{common.Address{}, nil},
			},
			{
				Values: gomonkey.Params{nil, errors.New("execution reverted: ERC721: invalid token ID")},
			},
		})
		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
		p.ApplyMethodReturn(caller, "NodeType", new(big.Int).SetUint64(1), nil)

		ch := make(chan *BlockProver, 10)
		err := listProver(ch, &prover.Prover{ProverCaller: *caller}, 0)
		r.NoError(err)
		res := <-ch
		r.Equal(res.BlockNumber, uint64(0))
		r.Equal(res.Provers[1].ID, uint64(1))
		r.Equal(*res.Provers[1].Paused, false)
	})
}

func TestWatchProver(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	c := make(chan time.Time, 10)
	p.ApplyFuncReturn(time.NewTicker, &time.Ticker{C: c})
	p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), nil)
	p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", []types.Log{}, nil)
	p.ApplyFuncReturn(processProverLogs, true)

	watchProver(nil, &ethclient.Client{}, &prover.Prover{}, time.Second, "", []common.Hash{{}, {}, {}, {}}, 0, 0)
	c <- time.Now()
	time.Sleep(20 * time.Millisecond)
	close(c)
}

func TestProcessProverLogs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	id := new(big.Int).SetUint64(1)
	filterer := &prover.ProverFilterer{}
	p.ApplyMethodReturn(filterer, "ParseOperatorSet", &prover.ProverOperatorSet{Id: id}, nil)
	p.ApplyMethodReturn(filterer, "ParseNodeTypeUpdated", &prover.ProverNodeTypeUpdated{Id: id, Typ: id}, nil)
	p.ApplyMethodReturn(filterer, "ParseProverPaused", &prover.ProverProverPaused{Id: id}, nil)
	p.ApplyMethodReturn(filterer, "ParseProverResumed", &prover.ProverProverResumed{Id: id}, nil)

	logs := []types.Log{
		{
			Topics:      []common.Hash{operatorSetTopicHash},
			BlockNumber: 100,
			TxIndex:     1,
		},
		{
			Topics:      []common.Hash{nodeTypeUpdatedTopicHash},
			BlockNumber: 99,
			TxIndex:     1,
		},
		{
			Topics:      []common.Hash{proverPausedTopicHash},
			BlockNumber: 100,
			TxIndex:     2,
		},
		{
			Topics:      []common.Hash{proverResumedTopicHash},
			BlockNumber: 101,
			TxIndex:     1,
		},
		{
			Topics:      []common.Hash{proverResumedTopicHash},
			BlockNumber: 101,
			TxIndex:     2,
		},
		{
			Topics:      []common.Hash{proverResumedTopicHash},
			BlockNumber: 98,
			TxIndex:     2,
		},
		{
			Topics:      []common.Hash{proverResumedTopicHash},
			BlockNumber: 98,
			TxIndex:     1,
		},
	}
	ps := make(chan *BlockProver, 10)
	processProverLogs(ps, logs, &prover.Prover{ProverFilterer: *filterer})
	r1 := <-ps
	r.Equal(r1.BlockNumber, uint64(98))
	r2 := <-ps
	r.Equal(r2.BlockNumber, uint64(99))
	r3 := <-ps
	r.Equal(r3.BlockNumber, uint64(100))
	r4 := <-ps
	r.Equal(r4.BlockNumber, uint64(101))
}
