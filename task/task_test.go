package task

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestTask_VerifySignature(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToDecodeTaskSignature", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, errors.New(t.Name()))

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("FailedToWriteBinary", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, errors.New(t.Name()))

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("FailedToWriteBinary2", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncSeq(binary.Write, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{errors.New(t.Name())},
				Times:  1,
			},
		})
		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("FailedToBufWriteString", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), errors.New(t.Name()))

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("FailedToBufWrite", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), errors.New(t.Name()))

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("FailedToRecoverPublicKey", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, errors.New(t.Name()))

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), t.Name())
	})
	t.Run("TaskSignatureUnmatched", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, nil)
		p.ApplyFuncReturn(bytes.Equal, false)

		tk := &Task{}
		r.ErrorContains(tk.VerifySignature(nil), "task signature unmatched")
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, nil)
		p.ApplyFuncReturn(bytes.Equal, true)

		tk := &Task{}
		r.NoError(tk.VerifySignature(nil))
	})
}

func TestStateLog_SignerAddress(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToDecodeTaskSignature", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToWriteBinary", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToWriteBinary2", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncSeq(binary.Write, []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{errors.New(t.Name())},
				Times:  1,
			},
		})

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWriteString", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWrite", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWrite2", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodSeq(buf, "Write", []gomonkey.OutputCell{
			{
				Values: gomonkey.Params{int(1), nil},
				Times:  1,
			},
			{
				Values: gomonkey.Params{int(1), errors.New(t.Name())},
				Times:  1,
			},
		})
		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToRecoverPublicKey", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToUnmarshalPublicKey", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, nil)
		p.ApplyFuncReturn(crypto.UnmarshalPubkey, nil, errors.New(t.Name()))

		tk := &Task{}
		ts := &StateLog{}
		_, err := ts.SignerAddress(tk)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		priKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		r.NoError(err)

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Ecrecover, nil, nil)
		p.ApplyFuncReturn(crypto.UnmarshalPubkey, &priKey.PublicKey, nil)

		tk := &Task{}
		ts := &StateLog{}
		_, err = ts.SignerAddress(tk)
		r.NoError(err)
	})
}

func TestState_String(t *testing.T) {
	r := require.New(t)
	r.Equal(StatePacked.String(), "packed")
	r.Equal(StateDispatched.String(), "dispatched")
	r.Equal(StateProved.String(), "proved")
	r.Equal(StateOutputted.String(), "outputted")
	r.Equal(StateFailed.String(), "failed")
	r.Equal(StateInvalid.String(), "invalid")
}
