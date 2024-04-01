package output

import (
	"testing"

	"github.com/machinefi/sprout/types"
	"github.com/stretchr/testify/require"
)

func Test_stdout_Output(t *testing.T) {
	r := require.New(t)
	o := &stdout{}
	_, err := o.Output(&types.Task{}, []byte("any"))
	r.NoError(err)
}
