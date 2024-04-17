package didvc_test

import (
	"net/http"
	"runtime"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/testutil"
)

func TestVerifyJWTCredential(t *testing.T) {
	if runtime.GOOS == `darwin` {
		return
	}

	r := require.New(t)
	p := gomonkey.NewPatches()

	t.Run("FailedToMarshalRequest", func(t *testing.T) {
		p = testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		_, err := didvc.VerifyJWTCredential("any", "any")
		r.ErrorContains(err, t.Name())
	})
	p = testutil.JsonMarshal(p, []byte("any"), nil)

	t.Run("FailedToDoPost", func(t *testing.T) {
		p = testutil.HttpPost(p, nil, errors.New(t.Name()))
		_, err := didvc.VerifyJWTCredential("any", "any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("UnexpectedRespondedStatus", func(t *testing.T) {
		p = testutil.HttpPost(p, &http.Response{
			StatusCode: 500,
			Body:       http.NoBody,
		}, nil)
		_, err := didvc.VerifyJWTCredential("any", "any")
		r.Error(err)
	})
	p = testutil.HttpPost(p, &http.Response{
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
	}, nil)

	t.Run("FailedToReadHttpBody", func(t *testing.T) {
		p = testutil.IoReadAll(p, []byte("any"), errors.New(t.Name()))
		_, err := didvc.VerifyJWTCredential("any", "any")
		r.ErrorContains(err, t.Name())
	})
	p = testutil.IoReadAll(p, []byte("any"), nil)

	t.Run("FailedToParseHttpBody", func(t *testing.T) {
		p = testutil.JsonUnmarshal(p, errors.New(t.Name()))
		_, err := didvc.VerifyJWTCredential("any", "any")
		r.ErrorContains(err, t.Name())
	})
	p = testutil.JsonUnmarshal(p, nil)
	_, err := didvc.VerifyJWTCredential("any", "any")
	r.NoError(err)

}
