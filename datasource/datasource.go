package datasource

import "github.com/machinefi/sprout/types"

type Datasource interface {
	Retrieve(nextTaskID uint64) (*types.Task, error)
}
