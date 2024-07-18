package output

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/task"
)

func Test_stdout_Output(t *testing.T) {
	r := require.New(t)
	o := &stdout{}
	_, err := o.Output(uint64(0), &task.Task{}, []byte("any"))
	r.NoError(err)
}
