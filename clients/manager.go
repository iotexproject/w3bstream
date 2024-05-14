package clients

import (
	_ "embed" // embed mock clients configuration
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

func NewManager(ioIDRegisterAddress, chainEndpoint, ioIDRegistryEndpoint string) (*Manager, error) {
	cli, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dail chain endpiont: %s", chainEndpoint)
	}

	instance, err := contracts.NewIoIDRegistry(common.HexToAddress(ioIDRegisterAddress), cli)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new ioIDRegistry")
	}

	return &Manager{
		mux:                  sync.Mutex{},
		cli:                  cli,
		ioIDRegistryInstance: instance,
		ioIDRegistryEndpoint: ioIDRegistryEndpoint,
		pool:                 make(map[string]*Client),
	}, nil
}

type Manager struct {
	mux                  sync.Mutex
	pool                 map[string]*Client
	cli                  *ethclient.Client
	ioIDRegistryInstance *contracts.IoIDRegistry
	ioIDRegistryEndpoint string
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

	// NOTE: this did is used for debug
	if id == "did:io:0x5b7902df415485c7e21334ca95ca94667278030e" {
		doc := []byte(`{"@context":["https://www.w3.org/ns/did/v1","https://w3id.org/security#keyAgreementMethod"],"id":"did:io:0x5b7902df415485c7e21334ca95ca94667278030e","authentication":["did:io:0x5b7902df415485c7e21334ca95ca94667278030e#Key-secp256k1-1"],"keyAgreement":["did:io:0x05b7e62aab18a50f778aa9ee7fd5cafccff184fa#Key-p256-2"],"verificationMethod":[{"id":"did:io:0x5b7902df415485c7e21334ca95ca94667278030e#Key-secp256k1-1","type":"JsonWebKey2020","controller":"did:io:0x5b7902df415485c7e21334ca95ca94667278030e","publicKeyJwk":{"crv":"secp256k1","x":"Gon1UzQ5nHiBHgfZEZB6Pm8e_jIaEqFZ7ST8u1X2KFY","y":"jRNCdq7MmINchIW_inBuqvrDq1PN0oDmzFuMFxGQj7E","kty":"EC","kid":"Key-secp256k1-1"}},{"id":"did:io:0x05b7e62aab18a50f778aa9ee7fd5cafccff184fa#Key-p256-2","type":"JsonWebKey2020","controller":"did:io:0x5b7902df415485c7e21334ca95ca94667278030e","publicKeyJwk":{"crv":"P-256","x":"BXcP9R0lFxRZ_RZV9LNBztqp3GOPLn2Iri0rq5ptq0Q","y":"jlxF9ZbSRdzmb8DlN33GG8M_AS4bf98-8pWcE6HyBgk","kty":"EC","kid":"Key-p256-2"}}]}`)
		jwk, err := ioconnect.NewJWKFromDoc(doc)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse did doc")
		}
		return &Client{jwk: jwk}, nil
	}

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
