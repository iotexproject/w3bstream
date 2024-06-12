package vm

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/task"
)

func TestHandler_Handle(t *testing.T) {
	r := require.New(t)

	h := NewHandler(
		map[Type]string{
			Risc0:  "any",
			Halo2:  "any",
			ZKwasm: "any",
		},
	)

	t.Run("MissingMessages", func(t *testing.T) {
		_, err := h.Handle(&task.Task{}, ZKwasm, "any", "any")
		r.Error(err)
	})

	t.Run("UnsupportedVMType", func(t *testing.T) {
		_, err := h.Handle(&task.Task{}, Type("other"), "any", "any")
		r.Error(err)
	})

	t.Run("FailedToNewVmInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(newInstance, nil, errors.New(t.Name()))
		_, err := h.Handle(&task.Task{}, ZKwasm, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToExecuteMessage", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(newInstance, &instance{}, nil)
		p.ApplyPrivateMethod(&instance{}, "release", func() {})
		p.ApplyPrivateMethod(&instance{}, "execute", func(context.Context, *task.Task) ([]byte, error) {
			return nil, errors.New(t.Name())
		})

		_, err := h.Handle(&task.Task{}, ZKwasm, "any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(newInstance, &instance{}, nil)
		p.ApplyPrivateMethod(&instance{}, "release", func() {})
		p.ApplyPrivateMethod(&instance{}, "execute", func(context.Context, *task.Task) ([]byte, error) {
			return []byte("any"), nil
		})
		p.ApplyFuncReturn(hex.DecodeString, []byte("any"), nil)

		_, err := h.Handle(&task.Task{}, ZKwasm, "any", "any")
		r.NoError(err)
	})
}
