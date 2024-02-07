package main

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

func initLogger() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func bindEnvConfig() {
	viper.MustBindEnv(ChainEndpoint)
	viper.MustBindEnv(HttpServiceEndpoint)
	viper.MustBindEnv(DatabaseDSN)
	viper.MustBindEnv(BootNodeMultiaddr)
	viper.MustBindEnv(IotexChainID)
	viper.MustBindEnv(ProjectContractAddress)
	viper.MustBindEnv(IPFSEndpoint)
	viper.MustBindEnv(DIDAuthServerEndpoint)

	viper.BindEnv(OperatorPrivateKey)
	viper.BindEnv(OperatorPrivateKeyED25519)
}
