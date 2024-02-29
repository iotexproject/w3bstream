package znodecontract

import (
	"reflect"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/machinefi/sprout/persistence/znode"
)

func PatchNewZnode(p *Patches, z *znode.Znode, err error) *Patches {
	return p.ApplyFunc(
		znode.NewZnode,
		func(_ common.Address, _ bind.ContractBackend) (*znode.Znode, error) {
			return z, err
		},
	)
}

var targetZnodeZnodeCaller = reflect.TypeOf(&znode.ZnodeCaller{})

func PatchZnodeCallerZnodes(p *Patches, did string, paused bool, err error) *Patches {
	return p.ApplyMethod(
		targetZnodeZnodeCaller,
		"Znodes",
		func(_ *znode.ZnodeCaller, opts *bind.CallOpts, arg0 uint64) (struct {
			Did    string
			Paused bool
		}, error) {
			return struct {
				Did    string
				Paused bool
			}{Did: did, Paused: paused}, err
		},
	)
}

func PatchZnodeCallerZnodesSeq(p *Patches, vs ...struct {
	Did    string
	Paused bool
	Err    error
}) *Patches {
	outputs := make([]OutputCell, 0)
	for _, v := range vs {
		outputs = append(outputs, OutputCell{
			Values: Params{
				struct {
					Did    string
					Paused bool
				}{
					Did:    v.Did,
					Paused: v.Paused,
				}, v.Err,
			},
			Times: 1,
		})
	}

	return p.ApplyMethodSeq(
		targetZnodeZnodeCaller,
		"Znodes",
		outputs,
	)
}
