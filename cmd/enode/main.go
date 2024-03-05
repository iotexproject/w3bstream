package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	solanaTypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/enode/api"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
	"github.com/pkg/errors"
)

func main() {
	initLogger()
	initConfig()

	pg, err := persistence.NewPostgres(viper.GetString(DatabaseDSN))
	if err != nil {
		log.Fatal(err)
	}

	_ = clients.NewManager()

	projectManager, err := project.NewManager(viper.GetString(ChainEndpoint), viper.GetString(ProjectContractAddress), viper.GetString(IPFSEndpoint))
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

func getENodeConfig() (*api.ENodeConfigInfoResp, error) {
	enodeConf := &api.ENodeConfigInfoResp{ProjectContractAddress: viper.GetString(ProjectContractAddress)}

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
