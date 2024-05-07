package api

import (
	"bytes"
	"encoding/json"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/sequencer/persistence"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestNewHttpServer(t *testing.T) {
	r := require.New(t)

	s := NewHttpServer(nil, uint(1), "", "", nil, true)
	r.Equal(uint(1), s.aggregationAmount)
}

func TestHttpServer_Run(t *testing.T) {
	r := require.New(t)

	s := &httpServer{
		engine: gin.Default(),
	}

	t.Run("FailedToRun", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gin.Engine{}, "Run", errors.New(t.Name()))

		err := s.Run("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gin.Engine{}, "Run", nil)

		err := s.Run("")
		r.NoError(err)
	})
}

func TestHttpServer_verifyToken(t *testing.T) {
	r := require.New(t)

	s := &httpServer{}
	t.Run("FailedToAuthorized", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		c.Request.Header.Set("authorization", "Bearer valid_token")

		p.ApplyFuncReturn(didvc.VerifyJWTCredential, "", errors.New(t.Name()))
		s.verifyToken(c)
		r.Equal(http.StatusUnauthorized, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
		c.Request.Header.Set("authorization", "Bearer valid_token")

		p.ApplyFuncReturn(didvc.VerifyJWTCredential, "clientID", nil)

		s.verifyToken(c)

		clientID, ok := didvc.ClientIDFrom(c.Request.Context())
		r.Equal(true, ok)
		r.Equal("clientID", clientID)
	})
}

func TestHttpServer_handleMessage(t *testing.T) {
	r := require.New(t)

	s := &httpServer{}

	t.Run("FailedToReadBody", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)

		p.ApplyFuncReturn(io.ReadAll, nil, errors.New(t.Name()))
		s.handleMessage(c)
		r.Equal(http.StatusInternalServerError, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("FailedToDecrypt", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`)))

		p.ApplyFuncReturn(io.ReadAll, []byte("body"), nil)
		p.ApplyFuncReturn(didvc.ClientIDFrom, "", true)
		p.ApplyMethodReturn(&ioconnect.JWK{}, "DecryptBySenderDID", nil, errors.New(t.Name()))
		s.handleMessage(c)
		r.Equal(http.StatusBadRequest, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("FailedToBindBody", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`)))

		p.ApplyFuncReturn(io.ReadAll, []byte("body"), nil)
		p.ApplyFuncReturn(didvc.ClientIDFrom, "", true)
		p.ApplyMethodReturn(&ioconnect.JWK{}, "DecryptBySenderDID", nil, nil)
		//p.ApplyFuncReturn(binding.JSON.BindBody, errors.New(t.Name()))
		s.handleMessage(c)
		r.Equal(http.StatusBadRequest, w.Code)
		//
		//actualResponse := &apitypes.ErrRsp{}
		//err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		//r.NoError(err)
		//r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("FailedToVerify", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`)))

		p.ApplyFuncReturn(io.ReadAll, []byte("body"), nil)
		p.ApplyFuncReturn(didvc.ClientIDFrom, "clientID", true)
		p.ApplyMethodReturn(&ioconnect.JWK{}, "DecryptBySenderDID", []byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`), nil)
		p.ApplyFuncReturn(clients.VerifyProjectPermissionByClientDID, errors.New(t.Name()))
		s.handleMessage(c)
		r.Equal(http.StatusUnauthorized, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("FailedToSave", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`)))

		p.ApplyFuncReturn(io.ReadAll, []byte("body"), nil)
		p.ApplyFuncReturn(didvc.ClientIDFrom, "clientID", true)
		p.ApplyMethodReturn(&ioconnect.JWK{}, "DecryptBySenderDID", []byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`), nil)
		p.ApplyFuncReturn(clients.VerifyProjectPermissionByClientDID, nil)
		p.ApplyMethodReturn(&persistence.Persistence{}, "Save", errors.New(t.Name()))
		s.handleMessage(c)
		r.Equal(http.StatusInternalServerError, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Contains(actualResponse.Error, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`)))

		p.ApplyFuncReturn(io.ReadAll, []byte("body"), nil)
		p.ApplyFuncReturn(didvc.ClientIDFrom, "clientID", true)
		p.ApplyMethodReturn(&ioconnect.JWK{}, "DecryptBySenderDID", []byte(`{"projectID": 123, "projectVersion": "v1", "data": "some data"}`), nil)
		p.ApplyFuncReturn(clients.VerifyProjectPermissionByClientDID, nil)
		p.ApplyMethodReturn(&persistence.Persistence{}, "Save", nil)
		s.handleMessage(c)
		r.Equal(http.StatusOK, w.Code)
	})
}
