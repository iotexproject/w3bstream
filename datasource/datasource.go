package datasource

import tasktype "github.com/iotexproject/w3bstream/task"

type Datasource interface {
	Retrieve(projectID, nextTaskID uint64) (*tasktype.Task, error)
}
