package apitypes

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewErrRsp(t *testing.T) {
	r := require.New(t)
	rsp := NewErrRsp(errors.New(t.Name()))
	r.Contains(rsp.Error, t.Name())
}
