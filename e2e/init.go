package e2e

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/e2e/utils"
	"github.com/iotexproject/w3bstream/service/apinode"
	apinodeconfig "github.com/iotexproject/w3bstream/service/apinode/config"
	apinodepersistence "github.com/iotexproject/w3bstream/service/apinode/persistence"
	"github.com/iotexproject/w3bstream/service/bootnode"
	"github.com/iotexproject/w3bstream/service/prover"
	proverconfig "github.com/iotexproject/w3bstream/service/prover/config"
	proverdb "github.com/iotexproject/w3bstream/service/prover/db"
	"github.com/iotexproject/w3bstream/service/sequencer"
	sequencerconfig "github.com/iotexproject/w3bstream/service/sequencer/config"
	sequencerdb "github.com/iotexproject/w3bstream/service/sequencer/db"
	"github.com/iotexproject/w3bstream/smartcontracts/go/fleetmanagement"
	"github.com/iotexproject/w3bstream/smartcontracts/go/mockproject"
	"github.com/iotexproject/w3bstream/smartcontracts/go/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/projectregistrar"
	"github.com/iotexproject/w3bstream/smartcontracts/go/router"
	"github.com/iotexproject/w3bstream/util/ipfs"
	libp2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
)

func bootNodeInit() (*bootnode.BootNode, error) {
	key, _, err := libp2pcrypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	node := bootnode.NewBootNode(bootnode.BootNodeConfig{
		PrivateKey:   key,
		Port:         8000,
		IoTeXChainID: 2,
	})

	return node, nil
}

func apiNodeInit(dbURI string, chainEndpoint string, bootnodeAddr string, taskManagerContractAddr string) (*apinode.APINode, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	cfg := apinodeconfig.Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9000",
		ProverServiceEndpoint:   "localhost:9002",
		AggregationAmount:       1,
		DatabaseDSN:             dbURI,
		PrvKey:                  "",
		BootNodeMultiAddr:       bootnodeAddr,
		IoTeXChainID:            2,
		ChainEndpoint:           chainEndpoint,
		BeginningBlockNumber:    0,
		TaskManagerContractAddr: taskManagerContractAddr,
	}

	db, err := apinodepersistence.NewPersistence(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	node := apinode.NewAPINode(&cfg, db, key)
	return node, nil
}

func sequencerInit(dbFile string, chainEndpoint string, bootnodeAddr string,
	contractDeployments *utils.ContractsDeployments,
) (*sequencer.Sequencer, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	db, err := sequencerdb.New(dbFile)
	if err != nil {
		return nil, err
	}

	cfg := &sequencerconfig.Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9001",
		BootNodeMultiAddr:       bootnodeAddr,
		IoTeXChainID:            2,
		ChainEndpoint:           chainEndpoint,
		ProverContractAddr:      contractDeployments.Prover,
		DaoContractAddr:         contractDeployments.Dao,
		MinterContractAddr:      contractDeployments.Minter,
		TaskManagerContractAddr: contractDeployments.TaskManager,
		BeginningBlockNumber:    0,
	}

	sq := sequencer.NewSequencer(cfg, db, key)
	return sq, nil
}

func proverInit(dbFile string, dbURI string, chainEndpoint string,
	contractDeployments *utils.ContractsDeployments,
) (*prover.Prover, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	db, err := proverdb.New(dbFile, crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return nil, err
	}

	cfg := &proverconfig.Config{
		LogLevel:        slog.LevelInfo,
		ServiceEndpoint: ":9002",
		VMEndpoints:     `{"1":"localhost:4001","2":"localhost:4002","3":"zkwasm:4001","4":"wasm:4001"}`,
		ChainEndpoint:   chainEndpoint,
		DatasourceDSN:   dbURI,
		// ProjectContractAddr:     projectContractAddr,
		// RouterContractAddr:      routerContractAddr,
		// TaskManagerContractAddr: taskManagerContractAddr,
		BeginningBlockNumber: 0,
	}

	prover := prover.NewProver(cfg, db, key)

	return prover, nil
}

func registerProject(t *testing.T, chainEndpoint string, ipfsURL string, projectFile string,
	contractDeployments *utils.ContractsDeployments, payerHex string) (*big.Int, error) {
	client, err := ethclient.Dial(chainEndpoint)
	require.NoError(t, err)
	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)
	payer, err := crypto.HexToECDSA(payerHex)
	require.NoError(t, err)

	// Register project with ioid
	mockProjectContract, err := mockproject.NewMockProject(
		common.HexToAddress(contractDeployments.MockProject), client)
	require.NoError(t, err)

	tOpts, err := bind.NewKeyedTransactorWithChainID(payer, chainID)
	require.NoError(t, err)

	tx, err := mockProjectContract.Register(tOpts)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	// Register project in w3bstream
	projectRegistrarContract, err := projectregistrar.NewProjectRegistrar(
		common.HexToAddress(contractDeployments.Registrar), client)
	require.NoError(t, err)
	projectID := big.NewInt(1)

	registerFee, err := projectRegistrarContract.RegistrationFee(nil)
	require.NoError(t, err)

	tOpts, err = bind.NewKeyedTransactorWithChainID(payer, chainID)
	require.NoError(t, err)
	tOpts.Value = registerFee

	tx, err = projectRegistrarContract.Register(tOpts, projectID)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	// Upload project file to IPFS and update project config
	ipfs := ipfs.NewIPFS(ipfsURL)
	content, err := os.ReadFile(projectFile)
	require.NoError(t, err)
	hash256 := sha256.Sum256(content)
	cid, err := ipfs.AddContent(content)
	require.NoError(t, err)
	projectFileURL := fmt.Sprintf("ipfs://%s/%s", ipfsURL, cid)
	wsProject, err := project.NewProject(common.HexToAddress(contractDeployments.WSProject), client)
	require.NoError(t, err)

	tOpts, err = bind.NewKeyedTransactorWithChainID(payer, chainID)
	require.NoError(t, err)

	tx, err = wsProject.UpdateConfig(tOpts, projectID, projectFileURL, hash256)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	tx, err = wsProject.Resume(tOpts, projectID)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	// Bind dapp to router
	router, err := router.NewRouter(common.HexToAddress(contractDeployments.Router), client)
	require.NoError(t, err)

	tx, err = router.BindDapp(tOpts, projectID, common.HexToAddress(contractDeployments.MockDapp))
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	return projectID, nil
}

func registerProver(t *testing.T, chainEndpoint string,
	contractDeployments *utils.ContractsDeployments, prover *ecdsa.PrivateKey) error {
	client, err := ethclient.Dial(chainEndpoint)
	require.NoError(t, err)
	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)

	fleetManagementContract, err := fleetmanagement.NewFleetManagement(
		common.HexToAddress(contractDeployments.FleetManagement), client)
	require.NoError(t, err)

	tOpts, err := bind.NewKeyedTransactorWithChainID(prover, chainID)
	require.NoError(t, err)

	tx, err := fleetManagementContract.Register(tOpts)
	require.NoError(t, err)

	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	return nil

}
