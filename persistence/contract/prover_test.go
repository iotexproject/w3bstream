package contract

import (
	"container/list"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

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

func TestBlockProvers(t *testing.T) {
	r := require.New(t)

	cp := &blockProvers{
		capacity: 2,
		blocks:   list.New(),
	}

	cp.add(&BlockProver{
		BlockNumber: 100,
		Provers: map[uint64]*Prover{
			1: {
				ID:        1,
				NodeTypes: 1,
			},
		},
	})
	ps := cp.provers(99)
	r.Equal(len(ps.Provers), 0)
	ps = cp.provers(100)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(1))

	cp.add(&BlockProver{
		BlockNumber: 101,
		Provers: map[uint64]*Prover{
			1: {
				ID:        1,
				NodeTypes: 2,
			},
		},
	})
	ps = cp.provers(100)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(1))
	ps = cp.provers(101)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(2))

	cp.add(&BlockProver{
		BlockNumber: 102,
		Provers: map[uint64]*Prover{
			1: {
				ID:        1,
				NodeTypes: 3,
			},
			2: {
				ID:        2,
				NodeTypes: 1,
			},
		},
	})
	ps = cp.provers(100)
	r.Equal(len(ps.Provers), 0)
	ps = cp.provers(101)
	r.Equal(len(ps.Provers), 1)
	ps = cp.provers(102)
	r.Equal(len(ps.Provers), 2)
	r.Equal(ps.Provers[1].NodeTypes, uint64(3))
	r.Equal(ps.Provers[2].NodeTypes, uint64(1))
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)

	cp.add(&BlockProver{
		BlockNumber: 103,
		Provers: map[uint64]*Prover{
			1: {
				ID:        1,
				NodeTypes: 2,
			},
		},
	})
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)
}

// func TestListProver(t *testing.T) {
// 	r := require.New(t)

// 	t.Run("FailedToGetOperator", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &prover.ProverCaller{}
// 		p.ApplyMethodReturn(caller, "Operator", nil, errors.New(t.Name()))

// 		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToGetPaused", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &prover.ProverCaller{}
// 		p.ApplyMethodReturn(caller, "Operator", common.Address{}, nil)
// 		p.ApplyMethodReturn(caller, "IsPaused", false, errors.New(t.Name()))

// 		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToGetNodeType", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &prover.ProverCaller{}
// 		p.ApplyMethodReturn(caller, "Operator", common.Address{}, nil)
// 		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
// 		p.ApplyMethodReturn(caller, "NodeType", nil, errors.New(t.Name()))

// 		err := listProver(nil, &prover.Prover{ProverCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &prover.ProverCaller{}
// 		p.ApplyMethodSeq(caller, "Operator", []gomonkey.OutputCell{
// 			{
// 				Values: gomonkey.Params{common.Address{}, nil},
// 			},
// 			{
// 				Values: gomonkey.Params{nil, errors.New("execution reverted: ERC721: invalid token ID")},
// 			},
// 		})
// 		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
// 		p.ApplyMethodReturn(caller, "NodeType", new(big.Int).SetUint64(1), nil)

// 		ch := make(chan *BlockProver, 10)
// 		err := listProver(ch, &prover.Prover{ProverCaller: *caller}, 0)
// 		r.NoError(err)
// 		res := <-ch
// 		r.Equal(res.BlockNumber, uint64(0))
// 		r.Equal(res.Provers[1].ID, uint64(1))
// 		r.Equal(*res.Provers[1].Paused, false)
// 	})
// }

// func TestProcessProverLogs(t *testing.T) {
// 	r := require.New(t)
// 	p := gomonkey.NewPatches()
// 	defer p.Reset()

// 	id := new(big.Int).SetUint64(1)
// 	filterer := &prover.ProverFilterer{}
// 	p.ApplyMethodReturn(filterer, "ParseOperatorSet", &prover.ProverOperatorSet{Id: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseNodeTypeUpdated", &prover.ProverNodeTypeUpdated{Id: id, Typ: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseProverPaused", &prover.ProverProverPaused{Id: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseProverResumed", &prover.ProverProverResumed{Id: id}, nil)

// 	logs := []types.Log{
// 		{
// 			Topics:      []common.Hash{operatorSetTopicHash},
// 			BlockNumber: 100,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{nodeTypeUpdatedTopicHash},
// 			BlockNumber: 99,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{proverPausedTopicHash},
// 			BlockNumber: 100,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{proverResumedTopicHash},
// 			BlockNumber: 101,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{proverResumedTopicHash},
// 			BlockNumber: 101,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{proverResumedTopicHash},
// 			BlockNumber: 98,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{proverResumedTopicHash},
// 			BlockNumber: 98,
// 			TxIndex:     1,
// 		},
// 	}
// 	ps := make(chan *BlockProver, 10)
// 	processProverLogs(ps, logs, &prover.Prover{ProverFilterer: *filterer})
// 	r1 := <-ps
// 	r.Equal(r1.BlockNumber, uint64(98))
// 	r2 := <-ps
// 	r.Equal(r2.BlockNumber, uint64(99))
// 	r3 := <-ps
// 	r.Equal(r3.BlockNumber, uint64(100))
// 	r4 := <-ps
// 	r.Equal(r4.BlockNumber, uint64(101))
// }
