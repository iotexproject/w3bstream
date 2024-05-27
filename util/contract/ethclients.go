package contract

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	// gEthClients global ethereum rpc clients pool
	gEthClients = map[string]Client{}
	// gMtxEthClients for thread-safe access gEthClients
	gMtxEthClients sync.Mutex
)

func NewEthClient(endpoint string) (Client, error) {
	gMtxEthClients.Lock()
	defer gMtxEthClients.Unlock()

	c, ok := gEthClients[endpoint]
	if ok {
		c.acquire()
		return c, nil
	}

	cc, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dail endpoint: %s", endpoint)
	}
	c = &client{
		Client:   cc,
		endpoint: endpoint,
	}
	c.acquire()
	gEthClients[endpoint] = c

	return c, nil
}

func ReleaseClient(c Client) {
	gMtxEthClients.Lock()
	defer gMtxEthClients.Unlock()

	c, ok := gEthClients[c.Endpoint()]
	if !ok {
		return
	}

	if c.release() <= 0 {
		c.close()
		delete(gEthClients, c.Endpoint())
	}
}

type Client interface {
	Endpoint() string
	bind.ContractBackend
	Counter
	close()
}

type client struct {
	*ethclient.Client
	endpoint string
	counter
}

func (c *client) Endpoint() string {
	return c.endpoint
}

func (c *client) close() {
	c.Client.Close()
}
