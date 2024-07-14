package main

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/vm"
)

var (
	projectCacheDir        = "./project_cache"
	localDBDir             = "./local_db"
	beginningBlockNumber   = uint64(26000000)
	chainEndpoint          = "https://babel-api.testnet.iotex.io"
	projectContractAddress = common.HexToAddress("0x2faBD8F8667158Ff8B0523f7BA8fC0CD0df3d0eA")
	proverContractAddress  = common.HexToAddress("0x0764e9c021F140d3A8CAb6EDd59904E584378D19")
	ipfsEndpoint           = "ipfs.mainnet.iotex.io"
	schedulerEpoch         = uint64(20)
)

var (
	halo2ProjectFile  = "./project_file/10001"
	risc0ProjectFile  = "./project_file/10000"
	zkwasmProjectFile = "./project_file/10001"
)

var (
	halo2MessageData  = "{\"private_a\": 3, \"private_b\": 4}"
	risc0MessageData  = "{\"private_input\":\"20\", \"public_input\":\"11,43\", \"receipt_type\":\"Stark\"}"
	zkwasmMessageData = "{\"private_input\": [1, 1] , \"public_input\": [] }"
)

func createProject() {
	cmd := exec.Command("./ioctl", "ioid", "register", uuid.New().String())
	o, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to register project", "error", err, "output", string(o))
		return
	}

	pidStr := strings.TrimSpace(strings.TrimPrefix(string(o), "Registerd ioID project id is"))
	slog.Info("currently project id", "project_id", pidStr)

	cmd = exec.Command("./ioctl", "ws", "project", "register", "--id", pidStr)
	o, err = cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to register project", "error", err, "output", string(o))
		return
	}

	switch rand.Intn(1) {
	case 0:
		cmd = exec.Command("./ioctl", "ws", "project", "update", "--id", pidStr, "--path", halo2ProjectFile)
		o, err = cmd.CombinedOutput()
		if err != nil {
			slog.Error("failed to update halo2 project", "project_id", pidStr, "error", err, "output", string(o))
			return
		}
		slog.Info("project halo2 config updated", "project_id", pidStr)
		// case 1: // will not create risc0 project
		// 	cmd = exec.Command("./ioctl", "ws", "project", "update", "--id", pidStr, "--path", risc0ProjectFile)
		// 	o, err = cmd.CombinedOutput()
		// 	if err != nil {
		// 		slog.Error("failed to update risc0 project", "project_id", pidStr, "error", err, "output", string(o))
		// 		return
		// 	}
		// 	slog.Info("project risc0 config updated", "project_id", pidStr)
		// case 2: // will not create zkwasm project
		// 	tx, err := projectInstance.UpdateConfig(opts, projectID, zkwasmProjectFileURI, zkwasmProjectFileHash)
		// 	if err != nil {
		// 		slog.Error("failed to update project config", "error", err)
		// 		return
		// 	}
		// 	slog.Info("project zkwasm config updated", "tx_hash", tx.Hash().Hex())
	}
	cmd = exec.Command("./ioctl", "ws", "project", "resume", "--id", pidStr)
	o, err = cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to resume project", "error", err, "output", string(o))
		return
	}
}

func updateProjectRequiredProver(contractPersistence *contract.Contract) {
	projects := contractPersistence.LatestProjects()
	provers := contractPersistence.LatestProvers()
	if len(provers) == 0 {
		slog.Error("no prover")
		return
	}

	index := rand.Intn(len(projects))
	project := projects[index]
	expectProvers := rand.Intn(len(provers)) + 1

	pidStr := strconv.FormatUint(project.ID, 10)

	cmd := exec.Command("./ioctl", "ws", "project", "attributes", "set", "--id", pidStr, "--key", "RequiredProverAmount", "--val", strconv.Itoa(expectProvers))
	var stdin bytes.Buffer
	stdin.Write([]byte("\n"))
	cmd.Stdin = &stdin
	o, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to update project attributes", "project_id", pidStr, "error", err, "output", string(o))
		return
	}
	slog.Info("project attributes setted", "project_id", project.ID, "prover_amount", expectProvers)
}

func sendMessage(contractPersistence *contract.Contract, projectManager *project.Manager, taskCount *atomic.Uint64) {
	projects := contractPersistence.LatestProjects()

	//slog.Info("send message", "project_len", len(projects))

	for i := 0; i < 3; i++ {
		index := rand.Intn(len(projects))
		project := projects[index]
		if project.Paused {
			continue
		}
		if project.Uri == "" {
			continue
		}
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

		response, err := http.Post("https://sprout-stress.w3bstream.com/message", "application/json", bytes.NewReader(j))
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
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	createProjectTicker := time.NewTicker(1 * time.Minute)
	go func() {
		for range createProjectTicker.C {
			createProject()
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
			updateProjectRequiredProver(contractPersistence)
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

	sendMessageTicker := time.NewTicker(1 * time.Second)
	go func() {
		for range sendMessageTicker.C {
			sendMessage(contractPersistence, projectManager, taskCount)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
