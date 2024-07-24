package vm

import (
	"encoding/hex"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/iotexproject/w3bstream/task"
)

func TestHandler_Handle(t *testing.T) {
	r := require.New(t)

	conn := &grpc.ClientConn{}
	h := &Handler{
		vmServerClients: map[Type]*grpc.ClientConn{
			Risc0:  conn,
			Halo2:  conn,
			ZKwasm: conn,
		},
	}
	t.Run("UnsupportedVMType", func(t *testing.T) {
		_, err := h.Handle(&task.Task{}, 1, "any", "any")
		r.Error(err)
	})
	t.Run("FailedToNewVmInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(create, errors.New(t.Name()))
		_, err := h.Handle(&task.Task{}, 1, "any", "any")
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToExecuteMessage", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(create, nil)
		p.ApplyFuncReturn(execute, nil, errors.New(t.Name()))

		_, err := h.Handle(&task.Task{}, 1, "any", "any")
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(create, nil)
		p.ApplyFuncReturn(execute, []byte("any"), nil)
		p.ApplyFuncReturn(hex.DecodeString, []byte("any"), nil)

		_, err := h.Handle(&task.Task{}, 1, "any", "any")
		r.NoError(err)
	})
}
