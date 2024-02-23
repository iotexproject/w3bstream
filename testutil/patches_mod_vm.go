package testutil

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

func VmHandlerHandle(p *Patches, err error) *Patches {
	var hander *vm.Handler
	return p.ApplyMethodFunc(
		reflect.TypeOf(hander),
		"Handle",
		func(msgs []*types.Message, vmtype types.VM, code string, expParam string) ([]byte, error) {
			return nil, err
		},
	)
}
