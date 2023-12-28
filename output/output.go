package output

import "github.com/machinefi/sprout/types"

type Output interface {
	Output(task *types.Task, proof []byte) (string, error)
}
