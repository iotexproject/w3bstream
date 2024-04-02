package project

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/project"
)

func ProjectConfigManagerGet(p *Patches, conf *project.ConfigData, err error) *Patches {
	var pm *project.ConfigManager
	return p.ApplyMethodFunc(
		reflect.TypeOf(pm),
		"Get",
		func(projectID uint64, version string) (*project.ConfigData, error) {
			return conf, err
		},
	)
}
