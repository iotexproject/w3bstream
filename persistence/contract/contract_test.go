package contract

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/smartcontracts/go/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/prover"
)

type mockCloser struct{}

func (m mockCloser) Close() error { return nil }

func TestContract_Project(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	t.Run("FailedToGet", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, pebble.ErrNotFound)

		project := c.Project(0, 0)
		r.Nil(project)
	})
	t.Run("FailedToUnmarshalBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", []byte("data"), nil, nil)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		project := c.Project(0, 0)
		r.Nil(project)
	})
	t.Run("FailedToCloseResultOfBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", []byte("data"), mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodReturn(mc, "Close", errors.New(t.Name()))

		project := c.Project(0, 0)
		r.Nil(project)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		bd := &block{
			blockProject: blockProject{
				Projects: map[uint64]*Project{},
			},
		}
		j, err := json.Marshal(bd)
		r.NoError(err)
		p.ApplyMethodReturn(c.db, "Get", j, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodReturn(mc, "Close", nil)

		project := c.Project(0, 0)
		r.Nil(project)
	})
}

func TestContract_LatestProject(t *testing.T) {
	r := require.New(t)
	c := &Contract{}
	t.Run("ProjectNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProjects", func() *blockProject { return nil })

		project := c.LatestProject(0)
		r.Nil(project)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProjects", func() *blockProject { return &blockProject{} })

		project := c.LatestProject(0)
		r.Nil(project)
	})
}

func TestContract_LatestProjects(t *testing.T) {
	r := require.New(t)
	c := &Contract{}
	t.Run("ProjectsNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProjects", func() *blockProject { return nil })

		projects := c.LatestProjects()
		r.Nil(projects)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProjects", func() *blockProject { return &blockProject{} })

		projects := c.LatestProjects()
		r.Equal(len(projects), 0)
	})
}

func TestContract_latestProjects(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, 10)
	t.Run("FailedToGetChainHeadData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, errors.New(t.Name()))

		projects := c.latestProjects()
		r.Nil(projects)
	})
	t.Run("FailedToCloseResultOfChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyMethodReturn(mc, "Close", errors.New(t.Name()))

		projects := c.latestProjects()
		r.Nil(projects)
	})
	t.Run("FailedToCloseResultOfChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodSeq(c.db, "Get", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{numberBytes, mc, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]byte{}, mc, pebble.ErrNotFound},
				Times:  1,
			},
		})

		projects := c.latestProjects()
		r.Nil(projects)
	})
	t.Run("FailedToUnmarshalBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		projects := c.latestProjects()
		r.Nil(projects)
	})
	t.Run("FailedToCloseResultOfBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodSeq(mc, "Close", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{errors.New(t.Name())},
				Times:  1,
			},
		})

		projects := c.latestProjects()
		r.Nil(projects)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)

		projects := c.latestProjects()
		r.NotNil(projects)
	})
}

func TestContract_Provers(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	t.Run("FailedToGet", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, pebble.ErrNotFound)

		prover := c.Provers(0)
		r.Nil(prover)
	})
	t.Run("FailedToUnmarshalBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", []byte("data"), nil, nil)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		prover := c.Provers(0)
		r.Nil(prover)
	})
	t.Run("FailedToCloseResultOfBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", []byte("data"), mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodReturn(mc, "Close", errors.New(t.Name()))

		prover := c.Provers(0)
		r.Nil(prover)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", []byte("data"), mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)

		prover := c.Provers(0)
		r.NotNil(prover)
	})
}

func TestContract_LatestProvers(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	t.Run("ProversNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProvers", func() *blockProver { return nil })

		provers := c.LatestProvers()
		r.Nil(provers)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		bp := &blockProver{Provers: map[uint64]*Prover{1: {OperatorAddress: common.Address{}}}}
		p.ApplyPrivateMethod(c, "latestProvers", func() *blockProver { return bp })

		provers := c.LatestProvers()
		r.NotNil(provers)
		r.Equal(len(provers), 1)
	})
}

func TestContract_Prover(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	t.Run("ProversNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "latestProvers", func() *blockProver { return nil })

		prover := c.Prover(common.Address{})
		r.Nil(prover)
	})
	t.Run("ProversNotFound", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		bp := &blockProver{Provers: map[uint64]*Prover{}}
		p.ApplyPrivateMethod(c, "latestProvers", func() *blockProver { return bp })

		prover := c.Prover(common.Address{})
		r.Nil(prover)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		bp := &blockProver{Provers: map[uint64]*Prover{1: {OperatorAddress: common.Address{}}}}
		p.ApplyPrivateMethod(c, "latestProvers", func() *blockProver { return bp })

		prover := c.Prover(common.Address{})
		r.NotNil(prover)
	})
}

func TestContract_latestProvers(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, 10)
	t.Run("FailedToGetChainHeadData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, errors.New(t.Name()))

		provers := c.latestProvers()
		r.Nil(provers)
	})
	t.Run("FailedToCloseResultOfChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyMethodReturn(mc, "Close", errors.New(t.Name()))

		provers := c.latestProvers()
		r.Nil(provers)
	})
	t.Run("FailedToCloseResultOfChainHead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodSeq(c.db, "Get", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{numberBytes, mc, nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{[]byte{}, mc, pebble.ErrNotFound},
				Times:  1,
			},
		})

		provers := c.latestProvers()
		r.Nil(provers)
	})
	t.Run("FailedToUnmarshalBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		provers := c.latestProvers()
		r.Nil(provers)
	})
	t.Run("FailedToCloseResultOfBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodSeq(mc, "Close", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{errors.New(t.Name())},
				Times:  1,
			},
		})

		provers := c.latestProvers()
		r.Nil(provers)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)

		provers := c.latestProvers()
		r.NotNil(provers)
	})
}

func TestContract_notifyProject(t *testing.T) {
	r := require.New(t)

	n := make(chan uint64, 10)
	c := &Contract{
		projectNotifications: []chan<- uint64{n},
	}
	c.notifyProject(&blockProjectDiff{
		diffs: map[uint64]*projectDiff{1: {id: 1}},
	})
	p := <-n
	r.Equal(uint64(1), p)
}

func TestContract_notifyChainHead(t *testing.T) {
	r := require.New(t)

	n := make(chan uint64, 10)
	c := &Contract{
		chainHeadNotifications: []chan<- uint64{n},
	}
	c.notifyChainHead(1)
	p := <-n
	r.Equal(uint64(1), p)
}

func TestContract_updateDB(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, 10)
	t.Run("FailedToGet", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, errors.New(t.Name()))

		err := c.updateDB(0, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnmarshalBlockData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		err := c.updateDB(0, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToCloseResultOfPreBlock", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		mc := mockCloser{}
		p.ApplyMethodReturn(c.db, "Get", numberBytes, mc, nil)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodReturn(mc, "Close", errors.New(t.Name()))

		err := c.updateDB(0, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToMarshal", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(c.db, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, errors.New(t.Name()))

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBatchSet", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		b := &pebble.Batch{}
		p.ApplyMethodReturn(c.db, "NewBatch", b)
		p.ApplyMethodReturn(b, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, nil)
		p.ApplyMethodReturn(b, "Set", errors.New(t.Name()))

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBatchSet2", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		b := &pebble.Batch{}
		p.ApplyMethodReturn(c.db, "NewBatch", b)
		p.ApplyMethodReturn(b, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, nil)
		p.ApplyMethodSeq(b, "Set", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{errors.New(t.Name())},
				Times:  1,
			},
		})

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBatchDel", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		b := &pebble.Batch{}
		p.ApplyMethodReturn(c.db, "NewBatch", b)
		p.ApplyMethodReturn(b, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, nil)
		p.ApplyMethodReturn(b, "Set", nil)
		p.ApplyMethodReturn(b, "Delete", errors.New(t.Name()))

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBatchCommit", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		b := &pebble.Batch{}
		p.ApplyMethodReturn(c.db, "NewBatch", b)
		p.ApplyMethodReturn(b, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, nil)
		p.ApplyMethodReturn(b, "Set", nil)
		p.ApplyMethodReturn(b, "Delete", nil)
		p.ApplyMethodReturn(b, "Commit", errors.New(t.Name()))

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		b := &pebble.Batch{}
		p.ApplyMethodReturn(c.db, "NewBatch", b)
		p.ApplyMethodReturn(b, "Get", nil, nil, pebble.ErrNotFound)
		p.ApplyFuncReturn(json.Marshal, nil, nil)
		p.ApplyMethodReturn(b, "Set", nil)
		p.ApplyMethodReturn(b, "Delete", nil)
		p.ApplyMethodReturn(b, "Commit", nil)

		err := c.updateDB(0, &blockProjectDiff{}, &blockProverDiff{})
		r.NoError(err)
	})
}

func TestContract_processLogs(t *testing.T) {
	r := require.New(t)
	c := &Contract{
		db: &pebble.DB{},
	}
	t.Run("FailedToProcessProjectLogs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "processProjectLogs", func([]types.Log) (map[uint64]*blockProjectDiff, error) { return nil, errors.New(t.Name()) })

		err := c.processLogs(0, 0, []types.Log{}, true)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToProcessProverLogs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "processProjectLogs", func([]types.Log) (map[uint64]*blockProjectDiff, error) { return nil, nil })
		p.ApplyPrivateMethod(c, "processProverLogs", func([]types.Log) (map[uint64]*blockProverDiff, error) { return nil, errors.New(t.Name()) })

		err := c.processLogs(0, 0, []types.Log{}, true)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUpdateDB", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "processProjectLogs", func([]types.Log) (map[uint64]*blockProjectDiff, error) {
			return map[uint64]*blockProjectDiff{0: {}}, nil
		})
		p.ApplyPrivateMethod(c, "processProverLogs", func([]types.Log) (map[uint64]*blockProverDiff, error) { return map[uint64]*blockProverDiff{}, nil })
		p.ApplyPrivateMethod(c, "updateDB", func(uint64, *blockProjectDiff, *blockProverDiff) error { return errors.New(t.Name()) })

		err := c.processLogs(0, 0, []types.Log{}, true)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(c, "processProjectLogs", func([]types.Log) (map[uint64]*blockProjectDiff, error) {
			return map[uint64]*blockProjectDiff{0: {}}, nil
		})
		p.ApplyPrivateMethod(c, "processProverLogs", func([]types.Log) (map[uint64]*blockProverDiff, error) { return map[uint64]*blockProverDiff{}, nil })
		p.ApplyPrivateMethod(c, "updateDB", func(uint64, *blockProjectDiff, *blockProverDiff) error { return nil })
		p.ApplyPrivateMethod(c, "notifyProject", func(*blockProjectDiff) {})

		err := c.processLogs(0, 0, []types.Log{}, true)
		r.NoError(err)
	})
}

func TestNew(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToDialChainEndpoint", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(nil, 1, 0, "", addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProjectContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(nil, 1, 0, "", addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProverContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(nil, 1, 0, "", addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToList", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, nil)
		p.ApplyPrivateMethod(&Contract{}, "list", func() (uint64, error) { return 0, errors.New(t.Name()) })

		addr := common.Address{}
		_, err := New(nil, 1, 0, "", addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, nil)
		p.ApplyPrivateMethod(&Contract{}, "list", func() (uint64, error) { return 0, nil })
		p.ApplyPrivateMethod(&Contract{}, "watch", func(uint64) {})

		addr := common.Address{}
		_, err := New(nil, 1, 0, "", addr, addr, nil, nil)
		time.Sleep(10 * time.Millisecond)

		r.NoError(err)
	})
}
