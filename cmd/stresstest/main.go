package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"log/slog"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	contractproject "github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/vm"
)

var (
	projectMinterPrivateKey string

	projectCacheDir        = "./project_cache"
	localDBDir             = "./local_db"
	beginningBlockNumber   = uint64(20000000)
	chainEndpoint          = "https://babel-api.testnet.iotex.io"
	projectContractAddress = common.HexToAddress("0xCBb7a80983Fd3405972F700101A82DB6304C6547")
	proverContractAddress  = common.HexToAddress("0x6B544a7603cead52AdfD99AA64B3d798083cc4CC")
	ipfsEndpoint           = "ipfs.mainnet.iotex.io"
	schedulerEpoch         = uint64(20)
)

var (
	halo2ProjectFileURI  = "ipfs://ipfs.mainnet.iotex.io/QmX3utHdGsPUmgouwnvCMQrE7RSJydpouEJQhqirbJepsf"
	halo2ProjectFileHash = common.HexToHash("0x5cb6a61aa744f9b3b8350e95ab04dda5f15876ea02b3bcdaf05912e6bfe17c0a")

	risc0ProjectFileURI  = "ipfs://ipfs.mainnet.iotex.io/QmZqwTe8yHSCLA61GkiR7qSzeMcBVSECiPjFEFZVT2R6tN"
	risc0ProjectFileHash = common.HexToHash("0x0a81cbc0463524fdcb48ecd9951b5d34ecd996ad1efbf3019465ada3dc6cc89a")

	zkwasmProjectFileURI  = "ipfs://ipfs.mainnet.iotex.io/QmWhzXMGrNusDY6axpmzV7sEmoUMdmZTa8ZqvKUK9N34af"
	zkwasmProjectFileHash = common.HexToHash("0x1e6779061e0004cb63366596f17e46272afd6ac6556933b1493dae3710c7ce88")
)

var (
	halo2MessageData  = "{\"private_a\": 3, \"private_b\": 4}"
	risc0MessageData  = "{\"private_input\":\"20\", \"public_input\":\"11,43\", \"receipt_type\":\"Stark\"}"
	zkwasmMessageData = "{\"private_input\": [1, 1] , \"public_input\": [] }"
)

func init() {
	flag.StringVar(&projectMinterPrivateKey, "projectMinterPrivateKey", "", "project minter private key")
}

func createProject(client *ethclient.Client, projectInstance *contractproject.Project, opts *bind.TransactOpts) {
	// tx, err := projectInstance.Mint(opts, opts.From) // TODO use new project contract logic
	// if err != nil {
	// 	slog.Error("failed to mint project", "error", err)
	// 	return
	// }
	// slog.Info("new project created", "tx_hash", tx.Hash().Hex())
	// for {
	// 	_, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
	// 	if err != nil {
	// 		slog.Error("failed to query tx hash", "error", err, "tx_hash", tx.Hash().Hex())
	// 		continue
	// 	}
	// 	if !isPending {
	// 		break
	// 	}
	// 	time.Sleep(3 * time.Second)
	// }
	// projectID, err := projectInstance.Count(nil)
	// if err != nil {
	// 	slog.Error("failed to query project id", "error", err)
	// 	return
	// }
	// slog.Info("new project created", "project_id", projectID.Uint64())

	// switch rand.Intn(2) {
	// case 0:
	// 	tx, err := projectInstance.UpdateConfig(opts, projectID, halo2ProjectFileURI, halo2ProjectFileHash)
	// 	if err != nil {
	// 		slog.Error("failed to update project config", "error", err)
	// 		return
	// 	}
	// 	slog.Info("project halo2 config updated", "tx_hash", tx.Hash().Hex())
	// case 1:
	// 	tx, err := projectInstance.UpdateConfig(opts, projectID, risc0ProjectFileURI, risc0ProjectFileHash)
	// 	if err != nil {
	// 		slog.Error("failed to update project config", "error", err)
	// 		return
	// 	}
	// 	slog.Info("project risc0 config updated", "tx_hash", tx.Hash().Hex())
	// case 2: // will not create zkwasm project
	// 	tx, err := projectInstance.UpdateConfig(opts, projectID, zkwasmProjectFileURI, zkwasmProjectFileHash)
	// 	if err != nil {
	// 		slog.Error("failed to update project config", "error", err)
	// 		return
	// 	}
	// 	slog.Info("project zkwasm config updated", "tx_hash", tx.Hash().Hex())
	// }
}

func updateProjectRequiredProver(contractPersistence *contract.Contract, projectInstance *contractproject.Project, opts *bind.TransactOpts) {
	projects := contractPersistence.LatestProjects()
	provers := contractPersistence.LatestProvers()
	if len(provers) == 0 {
		slog.Error("no prover")
		return
	}

	index := rand.Intn(len(projects))
	project := projects[index]
	expectProvers := rand.Intn(len(provers)) + 1

	tx, err := projectInstance.SetAttributes(opts, new(big.Int).SetUint64(project.ID), [][32]byte{contract.RequiredProverAmountHash}, [][]byte{[]byte(strconv.Itoa(expectProvers))})
	if err != nil {
		slog.Error("failed to set project attributes", "error", err)
		return
	}
	slog.Info("project attributes setted", "project_id", project.ID, "tx_hash", tx.Hash().Hex())
}

func sendMessage(contractPersistence *contract.Contract, projectManager *project.Manager, taskCount *atomic.Uint64) {
	projects := contractPersistence.LatestProjects()
	for i := 0; i < 3; i++ {
		index := rand.Intn(len(projects))
		project := projects[index]
		projectFile, err := projectManager.Project(project.ID)
		if err != nil {
			slog.Error("failed to get project data", "project_id", project.ID, "error", err)
			continue
		}
		defaultVer, err := projectFile.DefaultConfig()
		if err != nil {
			slog.Error("failed to get project default version", "project_id", project.ID, "error", err)
			continue
		}
		var data string
		switch defaultVer.VMType {
		case vm.Halo2:
			data = halo2MessageData
		case vm.Risc0:
			data = risc0MessageData
		case vm.ZKwasm:
			// data = zkwasmMessageData
			continue // skip zkwasm
		default:
			slog.Error("unsupported vm type", "project_id", project.ID, "vm_type", defaultVer.VMType)
			continue
		}
		req := &apitypes.HandleMessageReq{
			ProjectID:      project.ID,
			ProjectVersion: "0.1",
			Data:           data,
		}

		j, err := json.Marshal(req)
		if err != nil {
			slog.Error("failed to marshal json", "project_id", project.ID, "error", err)
			continue
		}

		response, err := http.Post("http://sprout-staging.w3bstream.com:9000/message", "application/json", bytes.NewReader(j))
		if err != nil {
			slog.Error("failed to send message", "project_id", project.ID, "error", err)
			continue
		}
		defer response.Body.Close()

		if !(response.StatusCode >= 200 && response.StatusCode <= 299) {
			slog.Error("failed to send message", "project_id", project.ID, "http_status_code", response.StatusCode)
			continue
		}
		taskCount.Add(1)
	}
}

func main() {
	flag.Parse()
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	priKey, err := crypto.HexToECDSA(projectMinterPrivateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse project minter private key"))
	}
	slog.Info("project minter address", "address", crypto.PubkeyToAddress(priKey.PublicKey).String())

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to dial chain endpoint"))
	}
	projectInstance, err := contractproject.NewProject(projectContractAddress, client)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project contract instance"))
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get chain id"))
	}

	opts := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(priKey.PublicKey),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewLondonSigner(chainID), priKey)
		},
	}

	createProjectTicker := time.NewTicker(10 * time.Minute)
	go func() {
		for range createProjectTicker.C {
			createProject(client, projectInstance, opts)
		}
	}()

	projectManagerNotification := make(chan uint64, 10)

	projectNotifications := []chan<- uint64{projectManagerNotification}

	db, err := pebble.Open(localDBDir, &pebble.Options{})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open pebble db"))
	}
	defer db.Close()

	contractPersistence, err := contract.New(db, schedulerEpoch, beginningBlockNumber, chainEndpoint, proverContractAddress,
		projectContractAddress, nil, projectNotifications)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
	}

	updateProjectRequiredProverTicker := time.NewTicker(1 * time.Hour)
	go func() {
		for range updateProjectRequiredProverTicker.C {
			updateProjectRequiredProver(contractPersistence, projectInstance, opts)
		}
	}()

	projectManager, err := project.NewManager(projectCacheDir, ipfsEndpoint, contractPersistence.LatestProject, projectManagerNotification)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	taskCount := &atomic.Uint64{}
	taskCountPrintTicker := time.NewTicker(1 * time.Minute)
	go func() {
		for range taskCountPrintTicker.C {
			slog.Info("number of tasks sent", "count", taskCount.Load())
		}
	}()

	sendMessageTicker := time.NewTicker(5 * time.Minute)
	go func() {
		for range sendMessageTicker.C {
			sendMessage(contractPersistence, projectManager, taskCount)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
