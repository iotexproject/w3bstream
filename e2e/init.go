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
	libp2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
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
	"github.com/iotexproject/w3bstream/smartcontracts/go/debits"
	"github.com/iotexproject/w3bstream/smartcontracts/go/mockerc20"
	"github.com/iotexproject/w3bstream/smartcontracts/go/mockproject"
	"github.com/iotexproject/w3bstream/smartcontracts/go/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/projectregistrar"
	"github.com/iotexproject/w3bstream/smartcontracts/go/projectreward"
	provercontract "github.com/iotexproject/w3bstream/smartcontracts/go/prover"
	"github.com/iotexproject/w3bstream/smartcontracts/go/router"
	"github.com/iotexproject/w3bstream/util/ipfs"
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

func apiNodeInit(dbURI string, chainEndpoint string, bootnodeAddr string, taskManagerContractAddr string) (*apinode.APINode, string, error) {
	cfg := apinodeconfig.Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9000",
		ProverServiceEndpoint:   "localhost:9002",
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
		return nil, "", err
	}

	node := apinode.NewAPINode(&cfg, db)
	return node, fmt.Sprintf("http://localhost%s", cfg.ServiceEndpoint), nil
}

func sequencerInit(dbURI string, dbFile string, chainEndpoint string, bootnodeAddr string,
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
		BootNodeMultiAddr:       bootnodeAddr,
		IoTeXChainID:            2,
		DatasourceDSN:           dbURI,
		ChainEndpoint:           chainEndpoint,
		ProverContractAddr:      contractDeployments.Prover,
		MinterContractAddr:      contractDeployments.Minter,
		TaskManagerContractAddr: contractDeployments.TaskManager,
		BeginningBlockNumber:    0,
	}

	sq := sequencer.NewSequencer(cfg, db, key)
	return sq, nil
}

func proverInit(dbURI string, dbFile string, chainEndpoint string, vmEndpoint string,
	contractDeployments *utils.ContractsDeployments,
) (*prover.Prover, *ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	db, err := proverdb.New(dbFile, crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return nil, nil, err
	}

	cfg := &proverconfig.Config{
		LogLevel:                slog.LevelInfo,
		ServiceEndpoint:         ":9002",
		VMEndpoints:             fmt.Sprintf("{\"1\":\"%s\"}", vmEndpoint),
		ChainEndpoint:           chainEndpoint,
		DatasourceDSN:           dbURI,
		ProjectContractAddr:     contractDeployments.WSProject,
		RouterContractAddr:      contractDeployments.Router,
		TaskManagerContractAddr: contractDeployments.TaskManager,
		BeginningBlockNumber:    0,
	}

	prover := prover.NewProver(cfg, db, key)

	return prover, key, nil
}

func registerProject(t *testing.T, chainEndpoint string, ipfsURL string, projectFile string,
	contractDeployments *utils.ContractsDeployments, projectOwner *ecdsa.PrivateKey) (*big.Int, error) {
	client, err := ethclient.Dial(chainEndpoint)
	require.NoError(t, err)
	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)

	// Register project with ioid
	mockProjectContract, err := mockproject.NewMockProject(
		common.HexToAddress(contractDeployments.MockProject), client)
	require.NoError(t, err)
	tOpts, err := bind.NewKeyedTransactorWithChainID(projectOwner, chainID)
	require.NoError(t, err)
	tx, err := mockProjectContract.Register(tOpts)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	newProjectID := big.NewInt(1)

	// Register project in w3bstream
	projectRegistrarContract, err := projectregistrar.NewProjectRegistrar(
		common.HexToAddress(contractDeployments.Registrar), client)
	require.NoError(t, err)
	registerFee, err := projectRegistrarContract.RegistrationFee(nil)
	require.NoError(t, err)
	tOpts, err = bind.NewKeyedTransactorWithChainID(projectOwner, chainID)
	require.NoError(t, err)
	tOpts.Value = registerFee
	tx, err = projectRegistrarContract.Register(tOpts, newProjectID)
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
	tOpts, err = bind.NewKeyedTransactorWithChainID(projectOwner, chainID)
	require.NoError(t, err)
	tx, err = wsProject.UpdateConfig(tOpts, newProjectID, projectFileURL, hash256)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	tx, err = wsProject.Resume(tOpts, newProjectID)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	// set reward for project
	var (
		rewardAmount = big.NewInt(100)
	)
	projectOwnerAddr := crypto.PubkeyToAddress(projectOwner.PublicKey)
	mockerc20Addr, _, mockerc20, err := mockerc20.DeployMockerc20(tOpts, client, "Mockerc20", "M20", projectOwnerAddr, big.NewInt(1e18))
	require.NoError(t, err)
	projectRewardContract, err := projectreward.NewProjectReward(
		common.HexToAddress(contractDeployments.ProjectReward), client)
	require.NoError(t, err)
	require.NoError(t, err)
	tx, err = projectRewardContract.SetReward(tOpts, newProjectID, rewardAmount)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	tx, err = projectRewardContract.SetRewardToken(tOpts, newProjectID, mockerc20Addr)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	tx, err = mockerc20.Approve(tOpts, common.HexToAddress(contractDeployments.Debits), rewardAmount)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	debitsContract, err := debits.NewDebits(common.HexToAddress(contractDeployments.Debits), client)
	require.NoError(t, err)
	tx, err = debitsContract.Deposit(tOpts, mockerc20Addr, rewardAmount)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	// Bind dapp to router
	router, err := router.NewRouter(common.HexToAddress(contractDeployments.Router), client)
	require.NoError(t, err)
	tx, err = router.BindDapp(tOpts, newProjectID, common.HexToAddress(contractDeployments.MockDapp))
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	return newProjectID, nil
}

func registerProver(t *testing.T, chainEndpoint string,
	contractDeployments *utils.ContractsDeployments, prover *ecdsa.PrivateKey) error {
	client, err := ethclient.Dial(chainEndpoint)
	require.NoError(t, err)
	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)

	proverContract, err := provercontract.NewProver(
		common.HexToAddress(contractDeployments.Prover), client)
	require.NoError(t, err)
	tOpts, err := bind.NewKeyedTransactorWithChainID(prover, chainID)
	require.NoError(t, err)
	tx, err := proverContract.Register(tOpts)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)
	tx, err = proverContract.SetRebateRatio(tOpts, 1000)
	require.NoError(t, err)
	_, err = utils.WaitForTransactionReceipt(client, tx.Hash())
	require.NoError(t, err)

	return nil
}
