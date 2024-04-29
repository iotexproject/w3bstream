package project

import (
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/persistence/contract"
)

func TestNewManager(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToNewCache", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(newCache, nil, errors.New(t.Name()))

		_, err := NewManager("cache", "", nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(newCache, nil, nil)

		_, err := NewManager("", "", nil, nil)
		r.NoError(err)
	})
}

func TestNewLocalManager(t *testing.T) {
	r := require.New(t)

	m := &Manager{}

	t.Run("FailedToLoadFromLocal", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(
			m,
			"loadFromLocal",
			func(projectFileDir string) error {
				return errors.New(t.Name())
			},
		)

		_, err := NewLocalManager("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(
			m,
			"loadFromLocal",
			func(projectFileDir string) error {
				return nil
			},
		)

		_, err := NewLocalManager("")
		r.NoError(err)
	})
}

func TestManager_Project(t *testing.T) {
	r := require.New(t)

	m := &Manager{}
	t.Run("FailedToLoad", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyPrivateMethod(&Manager{}, "load", func() (*Project, error) { return nil, errors.New(t.Name()) })
		_, err := m.Project(uint64(0))
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		m.projects.Store(uint64(0), &Project{})
		project, err := m.Project(uint64(0))
		r.NoError(err)
		r.Empty(project)
	})
}

func TestManager_ProjectIDs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	m := &Manager{}

	p.ApplyMethod(reflect.TypeOf(&sync.Map{}), "Range", func(m *sync.Map, f func(key, value interface{}) bool) {
		f(uint64(1), &Project{})
		f(uint64(2), &Project{})
	})

	ids := m.ProjectIDs()
	r.Equal(2, len(ids))
}

func TestManager_load(t *testing.T) {
	r := require.New(t)

	m := &Manager{
		contractProject: func(projectID uint64) *contract.Project {
			return &contract.Project{
				Uri:  "",
				Hash: common.Hash{},
			}
		},
		ipfsEndpoint: "https://ipfs.com",
		cache:        &cache{},
	}

	t.Run("NotExist", func(t *testing.T) {
		m1 := *m
		m1.contractProject = func(projectID uint64) *contract.Project {
			return nil
		}
		_, err := m1.load(uint64(0))
		r.ErrorContains(err, "the project not exist")
	})

	t.Run("FailedToGetRawData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&Meta{}, "FetchProjectRawData", nil, errors.New(t.Name()))
		_, err := m.load(uint64(0))
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToConvertProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&Meta{}, "FetchProjectRawData", []byte(""), nil)

		p.ApplyFuncReturn(convertProject, nil, errors.New(t.Name()))

		_, err := m.load(uint64(0))
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&Meta{}, "FetchProjectRawData", []byte(""), nil)
		p.ApplyFuncReturn(convertProject, &Project{}, nil)

		project, err := m.load(uint64(0))
		r.NoError(err)
		r.Empty(project)
	})
}

func TestManager_loadFromLocal(t *testing.T) {
	r := require.New(t)
	m := &Manager{}

	t.Run("FailedToReadDir", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(os.ReadDir, nil, errors.New(t.Name()))
		err := m.loadFromLocal("")
		r.ErrorContains(err, t.Name())
	})
}

func TestManager_watchProject(t *testing.T) {
	r := require.New(t)

	m := &Manager{}
	m.projects.Store(uint64(0), &Project{})

	projectNotification := make(chan *contract.Project)
	defer close(projectNotification)

	go m.watchProject(projectNotification)
	projectNotification <- &contract.Project{ID: 0}

	<-time.After(100 * time.Millisecond)

	_, ok := m.projects.Load(uint64(0))
	r.Equal(false, ok)
}
