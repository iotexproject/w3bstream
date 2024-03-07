package vm_test

import (
	"encoding/hex"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
	"github.com/machinefi/sprout/vm/server"
)

func TestHandler_Handle(t *testing.T) {
	r := require.New(t)

	h := vm.NewHandler(
		map[types.VM]string{
			types.VMRisc0:  "any",
			types.VMHalo2:  "any",
			types.VMZkwasm: "any",
		},
	)

	t.Run("MissingMessages", func(t *testing.T) {
		_, err := h.Handle(nil, types.VMZkwasm, "any", "any")
		r.Error(err)
	})

	t.Run("UnsupportedVMType", func(t *testing.T) {
		_, err := h.Handle([]*types.Message{{}}, types.VM("other"), "any", "any")
		r.Error(err)
	})

	t.Run("FailedToAcquireVmInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&server.Mgr{}, "Acquire", nil, errors.New(t.Name()))
		_, err := h.Handle([]*types.Message{{}}, types.VMZkwasm, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToExecuteMessage", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&server.Mgr{}, "Acquire", &server.Instance{}, nil)
		p = p.ApplyMethod(&server.Mgr{}, "Release", func(*server.Mgr, uint64, *server.Instance) {})
		p = p.ApplyMethodReturn(&server.Instance{}, "Execute", nil, errors.New(t.Name()))

		_, err := h.Handle([]*types.Message{{}}, types.VMZkwasm, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToHexDecode", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&server.Mgr{}, "Acquire", &server.Instance{}, nil)
		p = p.ApplyMethod(&server.Mgr{}, "Release", func(*server.Mgr, uint64, *server.Instance) {})
		p = p.ApplyMethodReturn(&server.Instance{}, "Execute", []byte("any"), nil)
		p = p.ApplyFuncReturn(hex.DecodeString, nil, errors.New(t.Name()))

		_, err := h.Handle([]*types.Message{{}}, types.VMZkwasm, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&server.Mgr{}, "Acquire", &server.Instance{}, nil)
		p = p.ApplyMethod(&server.Mgr{}, "Release", func(*server.Mgr, uint64, *server.Instance) {})
		p = p.ApplyMethodReturn(&server.Instance{}, "Execute", []byte("any"), nil)
		p = p.ApplyFuncReturn(hex.DecodeString, []byte("any"), nil)

		_, err := h.Handle([]*types.Message{{}}, types.VMZkwasm, "any", "any")
		r.NoError(err)
	})
}
