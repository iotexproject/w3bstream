package dispatcher

import (
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type NewDatasource func(datasourceURI string) (Datasource, error)

type Datasource interface {
	Retrieve(projectID, nextTaskID uint64) (*types.Task, error)
}

type Persistence interface {
	FetchNextTaskID(projectID uint64) (uint64, error)
	Create(s *types.TaskStateLog) error
}

type ProjectConfigManager interface {
	Get(projectID uint64, version string) (*project.ConfigData, error)
}

type Publisher interface {
	Publish(projectID uint64, data *p2p.Data) error
}
