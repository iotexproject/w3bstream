package datasource

import tasktype "github.com/machinefi/sprout/task"

type Datasource interface {
	Retrieve(projectID, nextTaskID uint64) (*tasktype.Task, error)
}
