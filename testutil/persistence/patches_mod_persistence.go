package persistence

import (
	"reflect"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/types"
)

func PersistencePostgresUpdateState(p *Patches, err error) *Patches {
	var pg *persistence.Postgres
	return p.ApplyMethodFunc(
		reflect.TypeOf(pg),
		"UpdateState",
		func(taskID string, state types.TaskState, comment string, createdAt time.Time) error {
			return err
		},
	)
}

func PersistencePostgresFetchByID(p *Patches, task *types.Task, err error) *Patches {
	var pg *persistence.Postgres
	return p.ApplyMethodFunc(
		reflect.TypeOf(pg),
		"FetchByID",
		func(taskID string) (*types.Task, error) {
			return task, err
		},
	)
}
