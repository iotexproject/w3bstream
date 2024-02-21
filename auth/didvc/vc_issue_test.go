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

func TestIssueCredential(t *testing.T) {
	if runtime.GOOS == `darwin` {
		return
	}

	r := require.New(t)
	p := gomonkey.NewPatches()

	req := &didvc.IssueCredentialReq{}
	t.Run("FailedToMarshalRequest", func(t *testing.T) {
		p = testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		_, err := didvc.IssueCredential("any", req, true)
		r.Error(err)
		r.Contains(err.Error(), t.Name())
	})
	p = testutil.JsonMarshal(p, []byte("any"), nil)

	t.Run("FailedToDoPost", func(t *testing.T) {
		p = testutil.HttpPost(p, &http.Response{}, errors.New(t.Name()))
		_, err := didvc.IssueCredential("any", req, true)
		r.Error(err)
		r.Contains(err.Error(), t.Name())
		r.NotNil(req.Options)
		r.Equal(req.Options.ProofFormat, didvc.ProofFormatJWT)
	})
	p = testutil.HttpPost(p, &http.Response{Body: http.NoBody}, nil)

	t.Run("FailedToReadHttpBody", func(t *testing.T) {
		p = testutil.IoReadAll(p, []byte("any"), errors.New(t.Name()))
		_, err := didvc.IssueCredential("any", req, true)
		r.Error(err)
		r.Contains(err.Error(), t.Name())
	})
	p = testutil.IoReadAll(p, []byte("any"), nil)

	t.Run("FailedToParseHttpBody", func(t *testing.T) {
		p = testutil.JsonUnmarshal(p, errors.New(t.Name()))
		_, err := didvc.IssueCredential("any", req, true)
		r.Error(err)
		r.Contains(err.Error(), t.Name())
	})
	p = testutil.JsonUnmarshal(p, nil)

	req.Credential.CredentialSubject = map[string]any{
		"id": "did:io:123123123123",
	}
	t.Run("FailedToCreateClientSession", func(t *testing.T) {
		p = testutil.ClientsCreateSession(p, errors.New(t.Name()))
		_, err := didvc.IssueCredential("any", req, true)
		r.Error(err)
		r.Contains(err.Error(), t.Name())
	})
	p = testutil.ClientsCreateSession(p, nil)
	_, err := didvc.IssueCredential("any", req, true)
	r.NoError(err)
}
