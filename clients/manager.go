package clients

import (
	_ "embed" // embed mock clients configuration
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/clients/contracts"
)

var (
	//go:embed clients
	mockClientsConfig []byte
)

type Client struct {
	ClientDID       string            `json:"clientDID"`
	Projects        []uint64          `json:"projects"`
	KeyAgreementKID string            `json:"keyAgreementKID"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	projects        map[uint64]struct{}
}

func (c *Client) init() {
	if c.Metadata == nil {
		c.Metadata = make(map[string]string)
	}
	if c.projects == nil {
		c.projects = make(map[uint64]struct{})
	}
	for _, v := range c.Projects {
		c.projects[v] = struct{}{}
	}
}

func (c *Client) HasProjectPermission(projectID uint64) bool {
	_, ok := c.projects[projectID]
	return ok
}

var manager *Manager

func NewManager(address string, chainEndpoint, ioIDRegistryEndpoint string) (*Manager, error) {
	if manager != nil {
		return manager, nil
	}

	cli, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dail chain endpiont: %s", chainEndpoint)
	}

	instance, err := contracts.NewIoIDRegistry(common.HexToAddress(address), cli)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new ioIDRegistry")
	}

	m := &Manager{
		mux:                  sync.Mutex{},
		cli:                  cli,
		ioIDRegistryInstance: instance,
		ioIDRegistryEndpoint: ioIDRegistryEndpoint,
		pool:                 make(map[string]*Client),
	}
	m.fillByMockClients()
	manager = m
	return manager, nil
}

type Manager struct {
	mux                  sync.Mutex
	pool                 map[string]*Client
	cli                  *ethclient.Client
	ioIDRegistryInstance *contracts.IoIDRegistry
	ioIDRegistryEndpoint string
}

func (mgr *Manager) clientByDIDFromCache(clientID string) (*Client, bool) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	c, ok := mgr.pool[clientID]
	return c, ok
}

func (mgr *Manager) ClientByDID(clientID string) *Client {
	c, ok := mgr.clientByDIDFromCache(clientID)
	if ok && c.KeyAgreementKID != "" {
		return c
	}

	kid, err := mgr.fetchClientFromContract(clientID)
	if err != nil {
		return c
	}
	c2 := c
	if c == nil {
		c2 = &Client{
			ClientDID: clientID,
			Projects:  make([]uint64, 0),
			Metadata:  make(map[string]string),
		}
		c2.init()
	}
	c2.KeyAgreementKID = kid

	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	mgr.pool[clientID] = c2
	return c2
}

func (mgr *Manager) fetchClientFromContract(clientID string) (string, error) {
	l := slog.With("client_id", clientID)

	clientAddress := strings.TrimPrefix(clientID, "did:io")
	l = l.With("client_address", clientAddress)

	uri, err := mgr.ioIDRegistryInstance.DocumentURI(nil, common.HexToAddress(clientAddress))
	if err != nil {
		err = errors.Wrap(err, "failed to read client document uri from contract")
		l.Error("read client document uri", "error", err)
		return "", err
	}
	l.With("ipfs_uri", uri)

	url := fmt.Sprintf("https://%s/cid/%s", mgr.ioIDRegistryEndpoint, uri)
	rsp, err := http.Get(url)
	if err != nil {
		err = errors.Wrap(err, "failed to fetch client did doc from io registry")
		l.Error("fetch client doc from io registry", "error", err)
		return "", err
	}

	defer rsp.Body.Close()
	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read response")
		l.Error("read document content from response", "error", err)
		return "", err
	}
	clientJWK, err := ioconnect.NewJWKFromDoc(content)
	if err != nil {
		err = errors.Wrap(err, "failed to parse did doc")
		l.Error("parse jwk from client doc", "error", err)
		return "", err
	}
	defer clientJWK.Destroy()

	return clientJWK.KeyAgreementKID(), nil
}

func (mgr *Manager) fetchProjectListByClientIDFromContract() {
	// TODO update client project binding list
}

func (mgr *Manager) AddClient(c *Client) {
	c.init()
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	mgr.pool[c.ClientDID] = c
}

func (mgr *Manager) fillByMockClients() {
	clients := make([]*Client, 0)
	if err := json.Unmarshal(mockClientsConfig, &clients); err != nil {
		panic(err)
	}
	for _, c := range clients {
		c.init()
		mgr.pool[c.ClientDID] = c
	}
}
