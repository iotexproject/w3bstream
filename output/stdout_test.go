package output

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stdout_Output(t *testing.T) {
	r := require.New(t)
	o := &stdout{}
	_, err := o.Output(1, nil, []byte("any"))
	r.NoError(err)
}
