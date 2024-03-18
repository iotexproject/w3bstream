package output

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
)

func Test_stdout_Output(t *testing.T) {
	r := require.New(t)
	o := &stdout{}
	_, err := o.Output(&types.Task{}, []byte("any"))
	r.NoError(err)
}
