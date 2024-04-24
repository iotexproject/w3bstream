package contract

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/util/hash"
)

func TestProject_Merge(t *testing.T) {
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
	np.Merge(diff)
	r.Equal(np, diff)
}

func TestBlockProject_Merge(t *testing.T) {
	r := require.New(t)

	np := &BlockProject{Projects: map[uint64]*Project{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &BlockProject{
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

	np.Merge(diff)
	r.Equal(np, diff)
}

func TestListAndWatchProject(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToDialChain", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := ListAndWatchProject("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProjectContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(project.NewProject, nil, errors.New(t.Name()))

		_, err := ListAndWatchProject("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToQueryChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(project.NewProject, &project.Project{}, nil)
		p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1), errors.New(t.Name()))

		_, err := ListAndWatchProject("", "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(project.NewProject, &project.Project{}, nil)
		p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1), nil)
		p.ApplyFuncReturn(listProject, nil)
		p.ApplyFuncReturn(watchProject)

		_, err := ListAndWatchProject("", "", 0)
		r.NoError(err)
	})
}

// func TestListProject(t *testing.T) {
// 	r := require.New(t)

// 	t.Run("FailedToGetConfig", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &project.ProjectCaller{}
// 		p.ApplyMethodReturn(caller, "Config", nil, errors.New(t.Name()))

// 		err := listProject(nil, &project.Project{ProjectCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToGetPaused", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &project.ProjectCaller{}
// 		p.ApplyMethodReturn(caller, "Config", project.W3bstreamProjectProjectConfig{}, nil)
// 		p.ApplyMethodReturn(caller, "IsPaused", false, errors.New(t.Name()))

// 		err := listProject(nil, &project.Project{ProjectCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToGetAttributes", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &project.ProjectCaller{}
// 		p.ApplyMethodReturn(caller, "Config", project.W3bstreamProjectProjectConfig{}, nil)
// 		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
// 		p.ApplyMethodReturn(caller, "Attributes", []byte{}, errors.New(t.Name()))

// 		err := listProject(nil, &project.Project{ProjectCaller: *caller}, 0)
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		caller := &project.ProjectCaller{}
// 		p.ApplyMethodSeq(caller, "Config", []gomonkey.OutputCell{
// 			{
// 				Values: gomonkey.Params{project.W3bstreamProjectProjectConfig{}, nil},
// 			},
// 			{
// 				Values: gomonkey.Params{nil, errors.New("execution reverted: ERC721: invalid token ID")},
// 			},
// 		})
// 		p.ApplyMethodReturn(caller, "IsPaused", false, nil)
// 		p.ApplyMethodReturn(caller, "Attributes", []byte{}, nil)

// 		ch := make(chan *BlockProject, 10)
// 		err := listProject(ch, &project.Project{ProjectCaller: *caller}, 0)
// 		r.NoError(err)
// 		res := <-ch
// 		r.Equal(res.BlockNumber, uint64(0))
// 		r.Equal(res.Projects[1].ID, uint64(1))
// 		r.Equal(*res.Projects[1].Paused, false)
// 	})
// }

func TestWatchProject(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	c := make(chan time.Time, 10)
	p.ApplyFuncReturn(time.NewTicker, &time.Ticker{C: c})
	p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), nil)
	p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", []types.Log{}, nil)
	p.ApplyFuncReturn(processProjectLogs, true)

	watchProject(nil, &ethclient.Client{}, &project.Project{}, time.Second, "", []common.Hash{{}, {}, {}, {}}, 0, 0)
	c <- time.Now()
	time.Sleep(20 * time.Millisecond)
	close(c)
}

// func TestProcessProjectLogs(t *testing.T) {
// 	r := require.New(t)
// 	p := gomonkey.NewPatches()
// 	defer p.Reset()

// 	id := new(big.Int).SetUint64(1)
// 	filterer := &project.ProjectFilterer{}
// 	p.ApplyMethodReturn(filterer, "ParseAttributeSet", &project.ProjectAttributeSet{ProjectId: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseProjectPaused", &project.ProjectProjectPaused{ProjectId: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseProjectResumed", &project.ProjectProjectResumed{ProjectId: id}, nil)
// 	p.ApplyMethodReturn(filterer, "ParseProjectConfigUpdated", &project.ProjectProjectConfigUpdated{ProjectId: id}, nil)

// 	logs := []types.Log{
// 		{
// 			Topics:      []common.Hash{attributeSetTopicHash},
// 			BlockNumber: 100,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{projectPausedTopicHash},
// 			BlockNumber: 99,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{projectResumedTopicHash},
// 			BlockNumber: 100,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{projectConfigUpdatedTopicHash},
// 			BlockNumber: 101,
// 			TxIndex:     1,
// 		},
// 		{
// 			Topics:      []common.Hash{projectConfigUpdatedTopicHash},
// 			BlockNumber: 101,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{projectConfigUpdatedTopicHash},
// 			BlockNumber: 98,
// 			TxIndex:     2,
// 		},
// 		{
// 			Topics:      []common.Hash{projectConfigUpdatedTopicHash},
// 			BlockNumber: 98,
// 			TxIndex:     1,
// 		},
// 	}
// 	ps := make(chan *BlockProject, 10)
// 	processProjectLogs(ps, logs, &project.Project{ProjectFilterer: *filterer})
// 	r1 := <-ps
// 	r.Equal(r1.BlockNumber, uint64(98))
// 	r2 := <-ps
// 	r.Equal(r2.BlockNumber, uint64(99))
// 	r3 := <-ps
// 	r.Equal(r3.BlockNumber, uint64(100))
// 	r4 := <-ps
// 	r.Equal(r4.BlockNumber, uint64(101))
// }
