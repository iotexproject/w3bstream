package datasource

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/task"
)

type Datasource interface {
	// If the provided taskID list is empty, an error will be reported
	// If the data for any task in the list cannot be retrieved, an error will be reported
	Retrieve(taskIDs []common.Hash) ([]*task.Task, error)
}
