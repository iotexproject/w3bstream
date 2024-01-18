package persistence

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/sprout/persistence/znode"
	"github.com/pkg/errors"
)

type ZNode struct {
	mux       sync.Mutex
	znodeDIDs map[string]bool

	contractAddress string
	chainEndpoint   string
}

func (z *ZNode) GetAll() []string {
	z.mux.Lock()
	defer z.mux.Unlock()

	dids := []string{}
	for d := range z.znodeDIDs {
		dids = append(dids, d)
	}
	return dids
}

// TODO monitor znode contract event
func NewZNode(chainEndpoint, contractAddress string) (*ZNode, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	}
	instance, err := znode.NewZnode(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "new znode contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}
	dids, err := instance.GetRunningZNodes(nil)
	if err != nil {
		return nil, errors.Wrap(err, "get running znode dids failed")
	}

	znodeDIDs := map[string]bool{}
	for _, d := range dids {
		znodeDIDs[d] = true
	}

	return &ZNode{
		znodeDIDs:       znodeDIDs,
		contractAddress: contractAddress,
		chainEndpoint:   chainEndpoint,
	}, nil
}
