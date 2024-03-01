package persistence

import (
	"sync"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/persistence/znode"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/testutil/znodecontract"
)

func PatchNewZnode(p *gomonkey.Patches, node *ZNode, err error) *gomonkey.Patches {
	return p.ApplyFunc(
		NewZNode,
		func(_, _ string) (*ZNode, error) {
			return node, err
		},
	)
}

func TestZnode(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("NewZnode", func(t *testing.T) {

		t.Run("FailedToDialEthClient", func(t *testing.T) {
			// mockey.PatchConvey(t.Name(), t, func() {
			// 	mockey.Mock(ethclient.Dial).Return(nil, errors.New(t.Name())).Build()
			// })
			p = testutil.EthClientDial(p, nil, errors.New(t.Name()))
			z, err := NewZNode("any", "any")
			r.Nil(z)
			r.ErrorContains(err, t.Name())
		})
		t.Run("FailedToCreateZnodeContractInstance", func(t *testing.T) {
			// mockey.PatchConvey(t.Name(), t, func() {
			// 	mockey.Mock(ethclient.Dial).Return(nil, nil).Build()
			// 	mockey.Mock(znode.NewZnode).Return(nil, errors.New(t.Name()))
			// })

			p = testutil.EthClientDial(p, nil, nil)
			p = znodecontract.PatchNewZnode(p, nil, errors.New(t.Name()))
			z, err := NewZNode("any", "any")
			r.Nil(z)
			r.ErrorContains(err, t.Name())
		})
		t.Run("Success", func(t *testing.T) {
			p = znodecontract.PatchNewZnode(p, &znode.Znode{
				ZnodeCaller:     znode.ZnodeCaller{},
				ZnodeFilterer:   znode.ZnodeFilterer{},
				ZnodeTransactor: znode.ZnodeTransactor{},
			}, nil)
			p = znodecontract.PatchZnodeCallerZnodesSeq(p,
				struct {
					Did    string
					Paused bool
					Err    error
				}{Err: errors.New(t.Name())},
				struct {
					Did    string
					Paused bool
					Err    error
				}{Did: "any", Err: nil},
				struct {
					Did    string
					Paused bool
					Err    error
				}{Did: "", Err: nil},
			)
			z, err := NewZNode("any", "any")
			r.NotNil(z)
			r.Nil(err)
		})
	})

	t.Run("Znode", func(t *testing.T) {
		p = PatchNewZnode(p, &ZNode{
			mux:             sync.Mutex{},
			znodeDIDs:       map[string]bool{"any1": true, "any2": true},
			contractAddress: "any",
			chainEndpoint:   "any",
		}, nil)

		z, err := NewZNode("any", "any")
		r.Nil(err)
		r.Len(z.GetAll(), 2)
	})
}
