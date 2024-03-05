package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	solanaTypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/enode/api"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

func main() {
	initLogger()
	initConfig()

	pg, err := persistence.NewPostgres(viper.GetString(DatabaseDSN))
	if err != nil {
		log.Fatal(err)
	}

	_ = clients.NewManager()

	projectManager, err := project.NewManager(viper.GetString(ChainEndpoint), viper.GetString(ProjectContractAddress), viper.GetString(ProjectFileDirectory), viper.GetString(IPFSEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := task.NewDispatcher(pg, projectManager, viper.GetString(BootNodeMultiaddr), viper.GetString(OperatorPrivateKey), viper.GetString(OperatorPrivateKeyED25519), viper.GetInt(IotexChainID))
	if err != nil {
		log.Fatal(err)
	}
	go dispatcher.Dispatch()

	go func() {
		if err := api.NewHttpServer(pg, viper.GetString(DIDAuthServerEndpoint), projectManager, getENodeConfig).Run(viper.GetString(HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func getENodeConfig() (*api.ENodeConfigResp, error) {
	enodeConf := &api.ENodeConfigResp{ProjectContractAddress: viper.GetString(ProjectContractAddress)}

	if len(viper.GetString(OperatorPrivateKey)) > 0 {
		pk := crypto.ToECDSAUnsafe(common.FromHex(viper.GetString(OperatorPrivateKey)))
		sender := crypto.PubkeyToAddress(pk.PublicKey)
		enodeConf.OperatorETHAddress = sender.String()
	}

	if len(viper.GetString(OperatorPrivateKeyED25519)) > 0 {
		wallet, err := solanaTypes.AccountFromHex(viper.GetString(OperatorPrivateKeyED25519))
		if err != nil {
			return nil, errors.Wrap(err, "get solana wallet failed")
		}
		enodeConf.OperatorSolanaAddress = wallet.PublicKey.String()
	}

	return enodeConf, nil
}
