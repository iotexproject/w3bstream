package contract

import (
	"container/list"
	"math/big"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/smartcontracts/go/multicall"
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
	np.merge(diff)
	r.Equal(np, diff)
}

func TestBlockProver_Merge(t *testing.T) {
	r := require.New(t)

	np := &blockProver{Provers: map[uint64]*Prover{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &blockProver{
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

	np.merge(diff)
	r.Equal(np, diff)
}

func TestBlockProvers_prover(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bps := &blockProvers{}
	p.ApplyPrivateMethod(bps, "provers", func(uint64) *blockProver {
		return &blockProver{}
	})
	addr := common.Address{}
	r.Nil(bps.prover(addr))

	p.ApplyPrivateMethod(bps, "provers", func(uint64) *blockProver {
		return &blockProver{
			Provers: map[uint64]*Prover{1: {
				OperatorAddress: addr,
			}},
		}
	})
	r.NotNil(bps.prover(addr))
}

func TestBlockProvers_provers(t *testing.T) {
	r := require.New(t)

	cp := &blockProvers{
		capacity: 2,
		blocks:   list.New(),
	}

	cp.add(&blockProver{
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

	cp.add(&blockProver{
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

	cp.add(&blockProver{
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

	ps = cp.provers(0)
	r.Equal(ps.Provers[1].NodeTypes, uint64(3))
	r.Equal(ps.Provers[2].NodeTypes, uint64(1))
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)

	cp.add(&blockProver{
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

func TestListProver(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToNewMultiCallContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToDecodeBlockNumberContractAbi", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyFuncReturn(abi.JSON, abi.ABI{}, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToDecodeProverContractAbi", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyFuncSeq(abi.JSON, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{abi.ABI{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{abi.ABI{}, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackBlockNumberCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodReturn(abi.ABI{}, "Pack", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProverOperatorCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProverIsPausedCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProverNodeTypeCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToMultiCall", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackBlockNumberResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", [][]byte{{}, {}, {}, {}}, nil)
		p.ApplyMethodReturn(abi.ABI{}, "Unpack", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProverOperatorResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", [][]byte{{}, []byte("1"), []byte("1"), []byte("1")}, nil)
		n := new(big.Int).SetUint64(1)
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		p.ApplyFuncSeq(abi.ConvertType, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProverIsPausedResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", [][]byte{{}, []byte("1"), []byte("1"), []byte("1")}, nil)
		n := new(big.Int).SetUint64(1)
		addr := common.Address{}
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&addr}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		p.ApplyFuncSeq(abi.ConvertType, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&addr},
				Times:  1,
			},
		})
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProverNodeTypeResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", [][]byte{{}, []byte("1"), []byte("1"), []byte("1")}, nil)
		n := new(big.Int).SetUint64(1)
		addr := common.Address{}
		paused := false
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&addr}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&paused}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		p.ApplyFuncSeq(abi.ConvertType, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&addr},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&paused},
				Times:  1,
			},
		})
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodSeq(caller, "MultiCall", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[][]byte{{}, []byte("1"), []byte("1"), []byte("1")}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[][]byte{{}, {}, {}, {}}, nil},
				Times:  1,
			},
		})
		n := new(big.Int).SetUint64(1)
		addr := common.Address{}
		paused := false
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  2,
			},
			{
				Values: gomonkey.Params{[]interface{}{&addr}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&paused}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
		})
		p.ApplyFuncSeq(abi.ConvertType, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&addr},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&paused},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
		})
		_, _, _, err := listProver(nil, addr, addr, addr)
		r.NoError(err)
	})
}

func TestProcessProverLogs(t *testing.T) {
	r := require.New(t)
	id := new(big.Int).SetUint64(1)
	filterer := &prover.ProverFilterer{}

	t.Run("FailedToParseProverOperatorSetEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseOperatorSet", &prover.ProverOperatorSet{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{operatorSetTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		err := processProverLogs(nil, logs, &prover.Prover{ProverFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverNodeTypeUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseNodeTypeUpdated", &prover.ProverNodeTypeUpdated{Id: id, Typ: id}, errors.New(t.Name()))
		logs := []types.Log{
			{
				Topics:      []common.Hash{nodeTypeUpdatedTopicHash},
				BlockNumber: 99,
				TxIndex:     1,
			},
		}

		err := processProverLogs(nil, logs, &prover.Prover{ProverFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverPausedUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProverPaused", &prover.ProverProverPaused{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{proverPausedTopicHash},
				BlockNumber: 100,
				TxIndex:     2,
			},
		}

		err := processProverLogs(nil, logs, &prover.Prover{ProverFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProverResumedUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProverResumed", &prover.ProverProverResumed{Id: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{proverResumedTopicHash},
				BlockNumber: 101,
				TxIndex:     1,
			},
		}

		err := processProverLogs(nil, logs, &prover.Prover{ProverFilterer: *filterer})
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

		ps := make(chan *blockProver, 10)
		processProverLogs(func(bp *blockProver) { ps <- bp }, logs, &prover.Prover{ProverFilterer: *filterer})
		r1 := <-ps
		r.Equal(r1.BlockNumber, uint64(98))
		r2 := <-ps
		r.Equal(r2.BlockNumber, uint64(99))
		r3 := <-ps
		r.Equal(r3.BlockNumber, uint64(100))
		r4 := <-ps
		r.Equal(r4.BlockNumber, uint64(101))
	})
}
