package clients

import (
	_ "embed" // embed mock clients configuration
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/clients/contracts"
)

func NewManager(
	projectClientContractAddress string,
	ioIDRegisterContractAddress string,
	ioIDRegistryServiceEndpoint string,
	chainEndpoint string,
) (*Manager, error) {
	cli, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dail chain endpiont: %s", chainEndpoint)
	}

	instanceIoID, err := contracts.NewIoIDRegistry(common.HexToAddress(ioIDRegisterContractAddress), cli)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new ioIDRegistry")
	}

	instanceProjectClient, err := contracts.NewProjectDevice(common.HexToAddress(projectClientContractAddress), cli)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new ioIDRegistry")
	}

	return &Manager{
		mux:                   sync.Mutex{},
		cli:                   cli,
		ioIDRegistryInstance:  instanceIoID,
		projectClientInstance: instanceProjectClient,
		ioIDRegistryEndpoint:  ioIDRegistryServiceEndpoint,
		pool:                  make(map[string]*Client),
	}, nil
}

type Manager struct {
	mux                   sync.Mutex
	pool                  map[string]*Client
	cli                   *ethclient.Client
	ioIDRegistryInstance  *contracts.IoIDRegistry
	projectClientInstance *contracts.ProjectDevice
	ioIDRegistryEndpoint  string
}

// clientByIoID get client from pool
func (mgr *Manager) clientByIoID(id string) *Client {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	return mgr.pool[id]
}

// ClientByIoID get client from pool, if not hit, try fetch client from contract
func (mgr *Manager) ClientByIoID(id string) *Client {
	c := mgr.clientByIoID(id)
	if c != nil {
		return c
	}

	c, err := mgr.fetchFromContract(id)
	if err != nil {
		slog.Error("fetch client", "error", err)
		return c
	}
	slog.Info("new client fetched from contract", "id", id, "did", c.DID(), "kid", c.KeyAgreementKID(), "client_doc", c.Doc())

	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	mgr.pool[c.DID()] = c
	return c
}

func (mgr *Manager) fetchFromContract(id string) (*Client, error) {
	address := strings.TrimPrefix(id, "did:io:")

	uri, err := mgr.ioIDRegistryInstance.DocumentURI(nil, common.HexToAddress(address))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read client document uri")
	}

	url := fmt.Sprintf("https://%s/cid/%s", mgr.ioIDRegistryEndpoint, uri)
	rsp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch client did doc from io registry")
	}

	defer rsp.Body.Close()
	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read io registry response")
	}
	jwk, err := ioconnect.NewJWKFromDoc(content)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse did doc")
	}
	return &Client{jwk: jwk}, nil
}

func (mgr *Manager) HasProjectPermission(clientID string, projectID uint64) (bool, error) {
	if c := mgr.clientByIoID(clientID); c == nil {
		return false, nil
	}

	clientAddress := strings.TrimPrefix(clientID, "did:io:")

	approved, err := mgr.projectClientInstance.Approved(nil, big.NewInt(int64(projectID)), common.HexToAddress(clientAddress))
	if err != nil {
		return false, errors.Wrap(err, "failed to read ProjectClient contract")
	}
	return approved, nil

}
