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
	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/util/hash"
)

func TestProject_merge(t *testing.T) {
	r := require.New(t)

	np := &Project{Attributes: map[common.Hash][]byte{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &Project{
		ID:          1,
		BlockNumber: 100,
		Paused:      &paused,
		Uri:         "uri",
		Hash:        hash,
		Attributes:  attr,
	}
	np.merge(diff)
	r.Equal(np, diff)
}

func TestBlockProject_merge(t *testing.T) {
	r := require.New(t)

	np := &blockProject{Projects: map[uint64]*Project{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &blockProject{
		BlockNumber: 100,
		Projects: map[uint64]*Project{
			1: {
				ID:          1,
				BlockNumber: 100,
				Paused:      &paused,
				Uri:         "uri",
				Hash:        hash,
				Attributes:  attr,
			},
		},
	}

	np.merge(diff)
	r.Equal(np, diff)
}

func TestBlockProjects_project(t *testing.T) {
	r := require.New(t)

	cp := &blockProjects{
		capacity: 2,
		blocks:   list.New(),
	}

	cp.add(&blockProject{
		BlockNumber: 100,
		Projects: map[uint64]*Project{
			1: {
				ID:  1,
				Uri: "uri",
			},
		},
	})
	p := cp.project(1, 99)
	r.Nil(p)
	p = cp.project(1, 100)
	r.Equal(p.ID, uint64(1))

	cp.add(&blockProject{
		BlockNumber: 101,
		Projects: map[uint64]*Project{
			1: {
				ID:  1,
				Uri: "uri1",
			},
		},
	})
	p = cp.project(1, 101)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri1")

	p = cp.project(1, 100)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri")

	cp.add(&blockProject{
		BlockNumber: 102,
		Projects: map[uint64]*Project{
			2: {
				ID:  2,
				Uri: "uri2",
			},
			1: {
				ID:  1,
				Uri: "uri2",
			},
		},
	})
	p = cp.project(2, 0)
	r.Equal(p.ID, uint64(2))
	r.Equal(p.Uri, "uri2")
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)

	p = cp.project(2, 102)
	r.Equal(p.ID, uint64(2))
	r.Equal(p.Uri, "uri2")
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)

	p = cp.project(1, 102)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri2")

	p = cp.project(1, 101)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri1")

	p = cp.project(1, 100)
	r.Nil(p)

	cp.add(&blockProject{
		BlockNumber: 105,
		Projects: map[uint64]*Project{
			1: {
				ID:  1,
				Uri: "uri1",
			},
		},
	})
	r.Equal(uint64(cp.blocks.Len()), cp.capacity)
}

func TestBlockProjects_projects(t *testing.T) {
	r := require.New(t)

	cp := &blockProjects{
		capacity: 2,
		blocks:   list.New(),
	}
	cp.add(&blockProject{
		BlockNumber: 100,
		Projects: map[uint64]*Project{
			1: {
				ID:  1,
				Uri: "uri",
			},
		},
	})
	res := cp.projects()
	r.Equal(res.BlockNumber, uint64(100))
}

func TestListProject(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToNewMultiCallContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToDecodeBlockNumberContractAbi", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyFuncReturn(abi.JSON, abi.ABI{}, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToDecodeProjectContractAbi", func(t *testing.T) {
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
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackBlockNumberCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodReturn(abi.ABI{}, "Pack", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackIsValidProjectCallData", func(t *testing.T) {
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
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProjectConfigCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  2,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProjectIsPausedCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  3,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProjectAttributesProverAmountCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  4,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToPackProjectAttributesVmTypeCallData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, nil, nil)
		p.ApplyMethodSeq(abi.ABI{}, "Pack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]byte{}, nil},
				Times:  5,
			},
			{
				Values: gomonkey.Params{nil, errors.New(t.Name())},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToMultiCall", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	multiCallRes := [][]byte{{}, {}, {}, {}, {}, {}}
	t.Run("FailedToUnpackBlockNumberResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		p.ApplyMethodReturn(abi.ABI{}, "Unpack", nil, errors.New(t.Name()))

		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackIsValidProjectResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
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
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProjectConfigResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		n := new(big.Int).SetUint64(1)
		isValidProject := true
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isValidProject}, nil},
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
				Values: gomonkey.Params{&isValidProject},
				Times:  1,
			},
		})
		addr := common.Address{}
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProverIsPausedResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		n := new(big.Int).SetUint64(1)
		isValidProject := true
		addr := common.Address{}
		conf := project.W3bstreamProjectProjectConfig{}
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isValidProject}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&conf}, nil},
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
				Values: gomonkey.Params{&isValidProject},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&conf},
				Times:  1,
			},
		})
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackProverAmtResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		n := new(big.Int).SetUint64(1)
		isValidProject := true
		addr := common.Address{}
		conf := project.W3bstreamProjectProjectConfig{}
		paused := false
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isValidProject}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&conf}, nil},
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
				Values: gomonkey.Params{&isValidProject},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&conf},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&paused},
				Times:  1,
			},
		})
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnpackVmTypeResult", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		n := new(big.Int).SetUint64(1)
		isValidProject := true
		addr := common.Address{}
		conf := project.W3bstreamProjectProjectConfig{}
		paused := false
		attr := []byte{}
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isValidProject}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&conf}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&paused}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&attr}, nil},
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
				Values: gomonkey.Params{&isValidProject},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&conf},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&paused},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&attr},
				Times:  1,
			},
		})
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(multicall.NewMulticall, &multicall.Multicall{}, nil)
		caller := &multicall.MulticallCaller{}
		p.ApplyMethodReturn(caller, "MultiCall", multiCallRes, nil)
		n := new(big.Int).SetUint64(1)
		isValidProject := true
		isInValidProject := false
		addr := common.Address{}
		conf := project.W3bstreamProjectProjectConfig{}
		paused := false
		attr := []byte{}
		p.ApplyMethodSeq(abi.ABI{}, "Unpack", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isValidProject}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&conf}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&paused}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&attr}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&attr}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&n}, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]interface{}{&isInValidProject}, nil},
				Times:  1,
			},
		})
		p.ApplyFuncSeq(abi.ConvertType, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&isValidProject},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&conf},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&paused},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&attr},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&attr},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&n},
				Times:  1,
			},
			{
				Values: gomonkey.Params{&isInValidProject},
				Times:  1,
			},
		})
		_, _, _, err := listProject(nil, addr, addr, addr)
		r.NoError(err)
	})
}

func TestProcessProjectLogs(t *testing.T) {
	r := require.New(t)
	id := new(big.Int).SetUint64(1)
	filterer := &project.ProjectFilterer{}

	t.Run("FailedToParseAttributeSetEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseAttributeSet", &project.ProjectAttributeSet{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{attributeSetTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		err := processProjectLogs(nil, logs, &project.Project{ProjectFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectPausedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectPaused", &project.ProjectProjectPaused{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectPausedTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		err := processProjectLogs(nil, logs, &project.Project{ProjectFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectResumedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectResumed", &project.ProjectProjectResumed{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectResumedTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		err := processProjectLogs(nil, logs, &project.Project{ProjectFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectConfigUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectConfigUpdated", &project.ProjectProjectConfigUpdated{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectConfigUpdatedTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		err := processProjectLogs(nil, logs, &project.Project{ProjectFilterer: *filterer})
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseAttributeSet", &project.ProjectAttributeSet{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectPaused", &project.ProjectProjectPaused{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectResumed", &project.ProjectProjectResumed{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectConfigUpdated", &project.ProjectProjectConfigUpdated{ProjectId: id}, nil)

		logs := []types.Log{
			{
				Topics:      []common.Hash{attributeSetTopicHash},
				BlockNumber: 100,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectPausedTopicHash},
				BlockNumber: 99,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectResumedTopicHash},
				BlockNumber: 100,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopicHash},
				BlockNumber: 101,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopicHash},
				BlockNumber: 101,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopicHash},
				BlockNumber: 98,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopicHash},
				BlockNumber: 98,
				TxIndex:     1,
			},
		}
		ps := make(chan *blockProject, 10)
		processProjectLogs(func(bp *blockProject) { ps <- bp }, logs, &project.Project{ProjectFilterer: *filterer})
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
