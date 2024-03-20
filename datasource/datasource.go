package datasource

import taskpkg "github.com/machinefi/sprout/task"

type Datasource interface {
	Retrieve(nextTaskID uint64) (*taskpkg.Task, error)
}
