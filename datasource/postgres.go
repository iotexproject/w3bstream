package datasource

import "github.com/machinefi/sprout/types"

type postgres struct{}

func (p *postgres) Retrieve(nextTaskID uint64) (*types.Task, error) {
	return nil, nil
}

func NewPostgres() (Datasource, error) {
	return nil, nil
}
