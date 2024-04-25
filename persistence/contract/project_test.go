package contract

import (
	"container/list"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

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

	np.Merge(diff)
	r.Equal(np, diff)
}

func TestBlockProjects(t *testing.T) {
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
