package contract

import (
	"bytes"
	"context"
	"encoding/hex"
	"reflect"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var (
	// gInstances global contract instances
	gInstances = map[string]*instance{}
	// gMtxInstances for thread-safe access gInstances
	gMtxInstances sync.Mutex
)

func NewInstance(name, address, endpoint string, abi abi.ABI) (Instance, error) {
	key := hex.EncodeToString([]byte(address + endpoint))

	gMtxInstances.Lock()
	defer gMtxInstances.Unlock()

	if i, ok := gInstances[key]; ok {
		i.acquire()
		return i, nil
	}

	backend, err := NewEthClient(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new contract instance")
	}

	_address := common.HexToAddress(address)
	i := &instance{
		name:     name,
		key:      key,
		backend:  backend,
		address:  _address,
		contract: bind.NewBoundContract(_address, abi, backend, backend, backend),
	}

	i.acquire()
	gInstances[key] = i
	return i, nil
}

func NewInstanceByABI(name, address, endpoint string, content []byte) (Instance, error) {
	_abi, err := abi.JSON(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	return NewInstance(name, address, endpoint, _abi)
}

func ReleaseInstance(i Instance) {
	gMtxInstances.Lock()
	defer gMtxInstances.Unlock()

	c, ok := gInstances[i.Key()]
	if !ok {
		return
	}

	if c.release() <= 0 {
		ReleaseClient(i.Client())
		delete(gInstances, i.Key())
	}
}

type Instance interface {
	Name() string
	Key() string
	Address() common.Address
	Client() Client
	Counter

	ReadContext(ctx context.Context, method string, args ...any) (any, error)
	Read(method string, args ...any) (any, error)
	ReadResultContext(ctx context.Context, method string, res any, args ...any) error
	ReadResult(method string, res any, args ...any) error
}

type instance struct {
	name     string
	key      string
	address  common.Address
	contract *bind.BoundContract
	backend  Client
	counter
}

func (i *instance) Name() string {
	return i.name
}

func (i *instance) Key() string {
	return i.key
}

func (i *instance) Address() common.Address {
	return i.address
}

func (i *instance) Client() Client {
	return i.backend
}

func (i *instance) Read(method string, args ...any) (any, error) {
	return i.ReadContext(context.Background(), method, args...)
}

func (i *instance) ReadContext(ctx context.Context, method string, args ...any) (any, error) {
	out := make([]any, 0)
	err := i.contract.Call(&bind.CallOpts{Context: ctx}, &out, method, args...)
	if err != nil {
		return nil, err
	}
	return out[0], nil
}

func (i *instance) ReadResult(method string, res any, args ...any) (err error) {
	return i.ReadResultContext(context.Background(), method, res, args...)
}

func (i *instance) ReadResultContext(ctx context.Context, method string, res any, args ...any) (err error) {
	rv, ok := res.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(res)
	}

	if !rv.IsValid() {
		return errors.Errorf("expect valid result, but got: (nil)")
	}

	rt := rv.Type()
	if rt.Kind() == reflect.Pointer {
		if rv.IsNil() && rv.CanSet() {
			rv.Set(reflect.New(rt.Elem()))
		}
		return i.ReadResultContext(ctx, method, rv.Elem(), args...)
	}
	if !rv.CanSet() {
		return errors.Errorf("expect result can be set, but got: %s", rt)
	}

	out, err := i.ReadContext(ctx, method, args...)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			err = errors.Errorf("%v", recover())
		}
	}()

	v := abi.ConvertType(out, reflect.New(rt).Interface())
	rv.Set(reflect.ValueOf(v).Elem())
	return err
}
