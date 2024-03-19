package output

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"io"
	"net/http"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tablelandnetwork/basin-cli/pkg/signing"

	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
)

func patchTextileDBPackData(p *Patches, data []byte, err error) *Patches {
	return p.ApplyPrivateMethod(&textileDB{}, "packData", func(*textileDB, []byte) ([]byte, error) {
		return data, err
	})
}

func patchTextileDBWrite(p *Patches, data string, err error) *Patches {
	return p.ApplyPrivateMethod(&textileDB{}, "write", func(*textileDB, []byte) (string, error) {
		return data, err
	})
}

func Test_textile_write(t *testing.T) {
	r := require.New(t)

	o := &textileDB{
		endpoint:  "any",
		secretKey: &ecdsa.PrivateKey{},
	}

	t.Run("FailedToSignData", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(signing.NewSigner, &signing.Signer{})
		p = p.ApplyMethodReturn(&signing.Signer{}, "SignBytes", nil, errors.New(t.Name()))

		txHash, err := o.write([]byte("any"))
		r.Equal(txHash, "")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToWriteTextileEvent", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(signing.NewSigner, &signing.Signer{})
		p = p.ApplyMethodReturn(&signing.Signer{}, "SignBytes", []byte("any"), nil)
		p = p.ApplyFuncReturn(writeTextileEvent, errors.New(t.Name()))

		txHash, err := o.write([]byte("any"))
		r.Equal(txHash, "")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(signing.NewSigner, &signing.Signer{})
		p = p.ApplyMethodReturn(&signing.Signer{}, "SignBytes", []byte("any"), nil)
		p = p.ApplyFuncReturn(writeTextileEvent, nil)

		txHash, err := o.write([]byte("any"))
		r.Equal(txHash, hex.EncodeToString([]byte("any")))
		r.NoError(err)
	})
}

func Test_textile_packData(t *testing.T) {
	r := require.New(t)

	o := &textileDB{
		endpoint:  "any",
		secretKey: &ecdsa.PrivateKey{},
	}

	t.Run("FailedToDecodeProof", func(t *testing.T) {
		proof := []byte("INVALID_HEX_DECODING")
		data, err := o.packData(proof)
		r.Nil(data)
		r.Error(err)
	})

	t.Run("MissingFieldJournalData", func(t *testing.T) {
		proof := []byte(hex.EncodeToString([]byte(`{}`)))
		data, err := o.packData(proof)
		r.Nil(data)
		r.Error(err)
	})

	t.Run("HasStarkJournalData", func(t *testing.T) {
		proof := []byte(hex.EncodeToString([]byte(`{"Stark":{"journal":{"bytes":[1]}}}`)))
		data, err := o.packData(proof)
		r.NotEmpty(data)
		r.NoError(err)
	})

	t.Run("HasSnarkJournalData", func(t *testing.T) {
		proof := []byte(hex.EncodeToString([]byte(`{"Snark":{"journal":[1]}}`)))
		data, err := o.packData(proof)
		r.NotEmpty(data)
		r.NoError(err)
	})

	t.Run("FailedToMarshalData", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testutil.JsonMarshal(p, nil, errors.New(t.Name()))

		proof := []byte(hex.EncodeToString([]byte(`{"Stark":{"journal":{"bytes":[1]}}}`)))
		data, err := o.packData(proof)
		r.Nil(data)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		proof := []byte(hex.EncodeToString([]byte(`{"Stark":{"journal":{"bytes":[1]}}}`)))
		data, err := o.packData(proof)
		r.NotEmpty(data)
		r.NoError(err)
	})
}

func Test_textile_Output(t *testing.T) {
	r := require.New(t)

	o := &textileDB{
		endpoint:  "any",
		secretKey: &ecdsa.PrivateKey{},
	}

	t.Run("FailedToPackData", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = patchTextileDBPackData(p, nil, errors.New(t.Name()))

		txHash, err := o.Output(&types.Task{}, []byte("any"))
		r.Empty(txHash)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToWrite", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = patchTextileDBPackData(p, []byte("any"), nil)
		p = patchTextileDBWrite(p, "", errors.New(t.Name()))

		txHash, err := o.Output(&types.Task{}, []byte("any"))
		r.Empty(txHash)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = patchTextileDBPackData(p, []byte("any"), nil)
		p = patchTextileDBWrite(p, "any", nil)

		txHash, err := o.Output(&types.Task{}, []byte("any"))
		r.NotEmpty(txHash)
		r.NoError(err)
	})
}

func Test_writeTextileEvent(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToNewHTTPRequest", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = testutil.HttpNewRequest(p, nil, errors.New(t.Name()))

		r.ErrorContains(writeTextileEvent("any", []byte("any")), t.Name())
	})

	t.Run("FailedToDoHTTPRequest", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = testutil.HttpNewRequest(p, &http.Request{Header: http.Header{}}, nil)
		p = p.ApplyMethodReturn(&http.Client{}, "Do", nil, errors.New(t.Name()))

		r.ErrorContains(writeTextileEvent("any", []byte("any")), t.Name())
	})

	t.Run("FailedToReadHttpResponse", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(http.NewRequest, &http.Request{
			Header: http.Header{},
		}, nil)
		p = p.ApplyMethodReturn(&http.Client{}, "Do", &http.Response{
			Body: io.NopCloser(bytes.NewReader(nil)),
		}, nil)
		p = testutil.IoReadAll(p, nil, errors.New(t.Name()))
		r.ErrorContains(writeTextileEvent("any", []byte("any")), t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(http.NewRequest, &http.Request{
			Header: http.Header{},
		}, nil)
		p = p.ApplyMethodReturn(&http.Client{}, "Do", &http.Response{
			Body: io.NopCloser(bytes.NewBufferString("any")),
		}, nil)
		p = testutil.IoReadAll(p, []byte("any"), nil)
		r.NoError(writeTextileEvent("any", []byte("any")))
	})
}
