package datasource

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/task"
)

type Datasource interface {
	Retrieve(projectID uint64, taskID common.Hash) (*task.Task, error)
}
