package contract_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	. "github.com/iotexproject/w3bstream/util/contract"
)

var (
	address = "0x06b3Fcda51e01EE96e8E8873F0302381c955Fddd"
	_abi    abi.ABI

	//go:embed testdata/test.json
	abiJSON []byte
)

func init() {
	var err error
	_abi, err = abi.JSON(bytes.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
}

func TestNewInstance(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToNewEthClient", func(t *testing.T) {
		i, err := NewInstance("any", "any", "invalid endpoint", abi.ABI{})
		r.Error(err)
		t.Log(err)
		r.Nil(i)
	})

	t.Run("Success", func(t *testing.T) {
		i1, err := NewInstance("TEST", address, endpoint, _abi)
		r.NoError(err)
		t.Log(i1.Name())
		t.Log(i1.Key())
		t.Log(i1.Address())
		t.Log(i1.RefCount())

		i2, err := NewInstance("TEST", address, endpoint, _abi)
		r.NoError(err)
		r.Equal(i1.Key(), i2.Key())
		r.Equal(i1.Name(), i2.Name())
		r.Equal(i1.Address(), i2.Address())
		r.Equal(i1.Client().Endpoint(), i2.Client().Endpoint())
		t.Log(i2.RefCount())
	})
}

func TestReleaseInstance(t *testing.T) {
	r := require.New(t)

	i, err := NewInstance("TEST", address, endpoint, _abi)
	r.NoError(err)
	for i.RefCount() > 0 {
		t.Log(i.RefCount())
		ReleaseInstance(i)
	}
	ReleaseInstance(i)
}

func TestInstance_ReadResult(t *testing.T) {
	r := require.New(t)

	i, err := NewInstance("TEST", address, endpoint, _abi)
	r.NoError(err)

	t.Run("InvalidValue", func(t *testing.T) {
		err = i.ReadResult(
			"documentURI",
			nil,
			common.HexToAddress("0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73"),
		)
		t.Log(err)
		r.ErrorContains(err, "expect valid result")
	})

	t.Run("CannotSetValue", func(t *testing.T) {
		err = i.ReadResult(
			"documentURI",
			1,
			common.HexToAddress("0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73"),
		)
		t.Log(err)
		r.ErrorContains(err, "expect result can be set")
	})

	t.Run("FailedToReadContract", func(t *testing.T) {
		err = i.ReadResult(
			"CannotFoundMethod",
			new(string),
			common.HexToAddress("0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73"),
		)
		t.Log(err)
		r.ErrorContains(err, "not found")
	})

	t.Run("ReflectPanic", func(t *testing.T) {
		res := new(int)
		err = i.ReadResult(
			"documentURI",
			res,
			common.HexToAddress("0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73"),
		)
		t.Log(err)
		r.Error(err)
	})

	t.Run("Success", func(t *testing.T) {
		res := (*string)(nil)
		err = i.ReadResult(
			"documentURI",
			&res,
			common.HexToAddress("0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73"),
		)
		r.NoError(err)
		t.Log(*res)
	})
}
