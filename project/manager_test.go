package project

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// func TestNewManager(t *testing.T) {
// 	r := require.New(t)

// 	t.Run("FailedToDialChain", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

// 		_, err := NewManager("", "", "", "", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToNewContracts", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)
// 		p.ApplyFuncReturn(project.NewProject, nil, errors.New(t.Name()))

// 		_, err := NewManager("", "", "", "", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToWatch", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)
// 		p.ApplyFuncReturn(project.NewProject, nil, nil)
// 		p.ApplyPrivateMethod(&Manager{}, "watchProjectContract", func() error { return errors.New(t.Name()) })

// 		_, err := NewManager("", "", "", "", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("FailedToNewCache", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)
// 		p.ApplyFuncReturn(project.NewProject, nil, nil)
// 		p.ApplyFuncReturn(newCache, nil, errors.New(t.Name()))

// 		_, err := NewManager("", "", "/cache", "", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)
// 		p.ApplyFuncReturn(project.NewProject, nil, nil)
// 		p.ApplyFuncReturn(newCache, nil, nil)
// 		p.ApplyPrivateMethod(&Manager{}, "watchProjectContract", func() error { return nil })

// 		_, err := NewManager("", "", "/cache", "", "")
// 		r.NoError(err)
// 	})
// }

func TestManager_Get(t *testing.T) {
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

// func TestManager_load(t *testing.T) {
// 	r := require.New(t)

// 	m := &Manager{
// 		instance: &project.Project{},
// 		cache:    &cache{},
// 	}

// 	t.Run("FailedToGetMeta", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyMethodReturn(&project.ProjectCaller{}, "Config", nil, errors.New(t.Name()))
// 		_, err := m.load(uint64(0))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("NotExist", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyMethodReturn(&project.ProjectCaller{}, "Config", project.W3bstreamProjectProjectConfig{}, nil)
// 		_, err := m.load(uint64(0))
// 		r.ErrorContains(err, "the project not exist")
// 	})

// 	t.Run("FailedToGetProject", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p = p.ApplyMethodReturn(
// 			&project.ProjectCaller{},
// 			"Config",
// 			project.W3bstreamProjectProjectConfig{Uri: "uri", Hash: [32]byte{1}},
// 			nil,
// 		)
// 		p.ApplyPrivateMethod(&cache{}, "get", func(projectID uint64, hash []byte) []byte { return []byte("") })
// 		p.ApplyMethodReturn(&Meta{}, "FetchProjectRawData", nil, errors.New(t.Name()))

// 		_, err := m.load(uint64(0))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("FailedToConvertProject", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyMethodReturn(
// 			&project.ProjectCaller{},
// 			"Config",
// 			project.W3bstreamProjectProjectConfig{Uri: "uri", Hash: [32]byte{1}},
// 			nil,
// 		)
// 		p.ApplyPrivateMethod(&cache{}, "get", func(projectID uint64, hash []byte) []byte { return []byte("data") })
// 		p.ApplyPrivateMethod(&cache{}, "set", func(projectID uint64, data []byte) {})
// 		p.ApplyFuncReturn(convertProject, nil, errors.New(t.Name()))

// 		_, err := m.load(uint64(0))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("Success", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyMethodReturn(
// 			&project.ProjectCaller{},
// 			"Config",
// 			project.W3bstreamProjectProjectConfig{Uri: "uri", Hash: [32]byte{1}},
// 			nil,
// 		)
// 		p.ApplyPrivateMethod(&cache{}, "get", func(projectID uint64, hash []byte) []byte { return []byte("data") })
// 		p.ApplyPrivateMethod(&cache{}, "set", func(projectID uint64, data []byte) {})
// 		p.ApplyFuncReturn(convertProject, &Project{}, nil)

// 		project, err := m.load(uint64(0))
// 		r.NoError(err)
// 		r.Empty(project)
// 	})
// }

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

// func TestManager_watchProjectContract(t *testing.T) {
// 	r := require.New(t)

// 	m := &Manager{}
// 	t.Run("FailedToWatch", func(t *testing.T) {
// 		p := gomonkey.NewPatches()
// 		defer p.Reset()

// 		p.ApplyFuncReturn(contract.ListAndWatchProject, nil, errors.New(t.Name()))

// 		err := m.watchProjectContract("", "")
// 		r.ErrorContains(err, t.Name())
// 	})
// }
