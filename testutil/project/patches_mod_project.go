package project

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/project"
)

func ProjectConfigManagerGet(p *Patches, conf *project.Config, err error) *Patches {
	var pm *project.Manager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"Get",
		func(projectID uint64, version string) (*project.Config, error) {
			return conf, err
		},
	)
}
