package project

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/project/contracts"
	wstypes "github.com/machinefi/sprout/types"
)

func TestNewManager(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToDialChain", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)

	t.Run("FailedToNewContracts", func(t *testing.T) {
		p = p.ApplyFuncReturn(contracts.NewContracts, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(contracts.NewContracts, nil, nil)

	t.Run("FailedToNewDefaultMonitor", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Manager{}, "fillProjectPool", func(string) {})
		p = p.ApplyFuncReturn(NewDefaultMonitor, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(NewDefaultMonitor, &Monitor{}, nil)

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Monitor{}, "run", func() {})
		p = p.ApplyMethodReturn(&Monitor{}, "MustEvents", make(chan *types.Log))
		p = p.ApplyPrivateMethod(&Manager{}, "watchProjectRegistrar", func(<-chan *types.Log, event.Subscription) {})

		_, err := NewManager("", "", "", "")
		r.NoError(err)
	})
}

type testSubscription struct {
	errChain chan error
}

func (s testSubscription) Err() <-chan error {
	return s.errChain
}

func (s testSubscription) Unsubscribe() {}

func TestManager_Get(t *testing.T) {
	r := require.New(t)

	t.Run("NotExist", func(t *testing.T) {
		m := &Manager{}
		_, err := m.Get(1, "0.1")
		r.ErrorContains(err, "project config not exist")
	})
	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			pool: map[key]*Config{getKey(1, "0.1"): {}},
		}
		_, err := m.Get(1, "0.1")
		r.NoError(err)
	})
}

func TestManager_Set(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			pool:       map[key]*Config{},
			projectIDs: map[uint64]bool{},
		}
		m.Set(1, "0.1", &Config{})
	})
}

func TestManager_GetAllProjectID(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
		}
		ids := m.GetAllProjectID()
		r.Equal(len(ids), 1)
		r.Equal(ids[0], uint64(1))
	})
}

func TestManager_GetNotify(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
			notify:     make(chan uint64, 1),
		}
		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		r.Equal(d, uint64(1))
	})
}

func TestManager_doProjectRegistrarWatch(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsFilterer{}, "ParseProjectUpserted", &contracts.ContractsProjectUpserted{ProjectId: 1}, nil)
		p = p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			notify:     make(chan uint64, 1),
			instance:   &contracts.Contracts{},
		}

		errChain := make(chan error)
		logChain := make(chan *types.Log, 1)
		logChain <- &types.Log{}

		m.doProjectRegistrarWatch(logChain, testSubscription{errChain})
		notify := m.GetNotify()
		d := <-notify
		r.Equal(d, uint64(1))
	})
}

func TestManager_fillProjectPoolFromContract(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToReadChain", func(t *testing.T) {
		outputs := []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{}, errors.New(t.Name())},
			},
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{}, nil},
			},
		}
		p = p.ApplyMethodSeq(&contracts.ContractsCaller{}, "Projects", outputs)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPoolFromContract()
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("Finished", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsCaller{}, "Projects", struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPoolFromContract()
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("FailedToGetConfig", func(t *testing.T) {
		outputs := []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{
					Uri:  "test",
					Hash: [32]byte{1},
				}, nil},
			},
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{}, nil},
			},
		}
		p = p.ApplyMethodSeq(&contracts.ContractsCaller{}, "Projects", outputs)
		p = p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", nil, errors.New(t.Name()))

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPoolFromContract()
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("Success", func(t *testing.T) {
		outputs := []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{
					Uri:  "test",
					Hash: [32]byte{1},
				}, nil},
			},
			{
				Values: gomonkey.Params{struct {
					Uri    string
					Hash   [32]byte
					Paused bool
				}{}, nil},
			},
		}
		p = p.ApplyMethodSeq(&contracts.ContractsCaller{}, "Projects", outputs)
		p = p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPoolFromContract()
		r.Equal(len(m.GetAllProjectID()), 1)
	})
}

type mockDirEntry struct {
	isDir bool
	name  string
}

func (m mockDirEntry) Name() string {
	return m.name
}

func (m mockDirEntry) IsDir() bool {
	return m.isDir
}

func (m mockDirEntry) Type() os.FileMode {
	return 0
}

func (m mockDirEntry) Info() (os.FileInfo, error) {
	return nil, nil
}

func TestManager_fillProjectPoolFromLocal(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	cs := []*Config{
		{
			Code:    "i am code",
			VMType:  wstypes.VMHalo2,
			Version: "0.1",
		},
	}
	jc, err := json.Marshal(cs)
	r.NoError(err)

	t.Run("ProjectFileDirParamIsEmpty", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("ReadDirNotExistError", func(t *testing.T) {
		p = p.ApplyFuncReturn(os.ReadDir, nil, os.ErrNotExist)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("FailedToReadDir", func(t *testing.T) {
		p = p.ApplyFuncReturn(os.ReadDir, nil, errors.New(t.Name()))

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("FailedToReadFile", func(t *testing.T) {
		p = p.ApplyFuncReturn(os.ReadDir, []os.DirEntry{mockDirEntry{isDir: true}, mockDirEntry{}}, nil)
		p = p.ApplyFuncReturn(os.ReadFile, nil, errors.New(t.Name()))

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("InvalidFileName", func(t *testing.T) {
		p = p.ApplyFuncReturn(os.ReadFile, nil, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyFuncReturn(os.ReadDir, []os.DirEntry{mockDirEntry{isDir: true}, mockDirEntry{name: "1"}}, nil)
		p = p.ApplyFuncReturn(os.ReadFile, jc, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 1)
	})
	t.Run("FailedToUnmarshalJson", func(t *testing.T) {
		p = p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
		}
		m.fillProjectPoolFromLocal("test")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
}

func TestManager_fillProjectPool(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Manager{}, "fillProjectPoolFromContract", func() {})
		p = p.ApplyPrivateMethod(&Manager{}, "fillProjectPoolFromLocal", func(string) {})

		m := &Manager{
			projectIDs: map[uint64]bool{},
		}
		m.fillProjectPool("")
		r.Equal(len(m.GetAllProjectID()), 0)
	})
}
