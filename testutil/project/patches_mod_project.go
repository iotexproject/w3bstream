package project

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/project"
)

func ProjectManagerGetAllProjectID(p *Patches, ids []uint64) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"GetAllProjectID",
		func() []uint64 {
			return ids
		},
	)
}

func ProjectManagerGetNotify(p *Patches, c <-chan uint64) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"GetNotify",
		func() <-chan uint64 {
			return c
		},
	)
}

func ProjectManagerGet(p *Patches, conf *project.Project, err error) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"Get",
		func(projectID uint64, version string) (*project.Project, error) {
			return conf, err
		},
	)
}

func ProjectConfigGetOutput(p *Patches, ot output.Output, err error) *Patches {
	var config *project.Config
	return p.ApplyMethodFunc(
		reflect.TypeOf(config),
		"GetOutput",
		func(privateKeyECDSA, privateKeyED25519 string) (output.Output, error) {
			return ot, err
		},
	)
}
