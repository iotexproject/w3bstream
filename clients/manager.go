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

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/iotexproject/w3bstream/util/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"
)

var (
	//go:embed contracts/ioIDRegistry.json
	abiIoIDRegistry []byte
	//go:embed contracts/ProjectDevice.json
	abiProjectClient []byte
	//go:embed contracts/W3bstreamProject.json
	abiW3bstreamProject []byte
)

func NewManager(
	projectClientContractAddress string,
	ioIDRegisterContractAddress string,
	w3bstreamProjectContractAddress string,
	ioIDRegistryServiceEndpoint string,
	chainEndpoint string,
) (*Manager, error) {
	manager := &Manager{
		mux:                  sync.Mutex{},
		pool:                 make(map[string]*Client),
		ioIDRegistryEndpoint: ioIDRegistryServiceEndpoint,
	}

	{
		name := "IoIDRegistry"
		instance, err := contract.NewInstanceByABI(name, ioIDRegisterContractAddress, chainEndpoint, abiIoIDRegistry)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new contract instance: %s", name)
		}
		manager.ioIDRegistryInstance = instance
	}

	{
		name := "ProjectClient"
		instance, err := contract.NewInstanceByABI(name, projectClientContractAddress, chainEndpoint, abiProjectClient)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new contract instance: %s", name)
		}
		manager.projectClientInstance = instance
	}

	{
		name := "W3bstreamProject"
		instance, err := contract.NewInstanceByABI(name, w3bstreamProjectContractAddress, chainEndpoint, abiW3bstreamProject)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new contract instance: %s", name)
		}
		manager.w3bsteramProjectInstance = instance
	}

	return manager, nil
}

type Manager struct {
	mux                      sync.Mutex
	pool                     map[string]*Client
	ioIDRegistryInstance     contract.Instance
	projectClientInstance    contract.Instance
	w3bsteramProjectInstance contract.Instance
	ioIDRegistryEndpoint     string
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
	var (
		address = common.HexToAddress(strings.TrimPrefix(id, "did:io:"))
		uri     string
	)

	if err := mgr.ioIDRegistryInstance.ReadResult("documentURI", &uri, address); err != nil {
		return nil, errors.Wrapf(err, "failed to read client document uri: %s %s", id, address.String())
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

	var (
		_projectID = big.NewInt(int64(projectID))
		attrVal    = make([]byte, 0)
		attrKey    = crypto.Keccak256Hash([]byte("key")) // todo replace with correct key name
		instance   = mgr.projectClientInstance           // use default contract
	)
	if err := mgr.w3bsteramProjectInstance.ReadResult("attribute", &attrVal, _projectID, attrKey); err == nil && len(attrVal) > 0 {
		slog.Info("get project attribute", "key", "key", "value", string(attrVal))
		name := fmt.Sprintf("ProjectClientFor%d", projectID)
		// todo should try acquire exists contract instance by name
		// todo should parse attribute value to project assigned contract meta
		projectAddress := ""
		projectChain := ""
		projectABI := "" // should have same interface for query is client is approved
		if _instance, err := contract.NewInstanceByABI(name, projectAddress, projectChain, []byte(projectABI)); err == nil {
			instance = _instance // use project assigned contract
		}
	}

	var (
		address  = common.HexToAddress(strings.TrimPrefix(clientID, "did:io:"))
		approved bool
	)

	if err := instance.ReadResult("approved", &approved, _projectID, address); err != nil {
		return false, errors.Wrapf(err, "failed to read ProjectClient contract: %s", instance.Address())
	}
	return approved, nil
}
