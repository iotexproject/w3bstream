package main

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/enode/constant"
)

func initLogger() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func initConfig() {
	viper.SetDefault(constant.ChainEndpoint, "https://babel-api.testnet.iotex.io")
	viper.SetDefault(constant.HttpServiceEndpoint, ":9000")
	viper.SetDefault(constant.DatabaseDSN, "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable")
	viper.SetDefault(constant.BootNodeMultiaddr, "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o")
	viper.SetDefault(constant.IotexChainID, 2)
	viper.SetDefault(constant.ProjectContractAddress, "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26")
	viper.SetDefault(constant.IPFSEndpoint, "ipfs.mainnet.iotex.io")
	viper.SetDefault(constant.DIDAuthServerEndpoint, "didkit:9999")

	viper.MustBindEnv(constant.ChainEndpoint)
	viper.MustBindEnv(constant.HttpServiceEndpoint)
	viper.MustBindEnv(constant.DatabaseDSN)
	viper.MustBindEnv(constant.BootNodeMultiaddr)
	viper.MustBindEnv(constant.IotexChainID)
	viper.MustBindEnv(constant.ProjectContractAddress)
	viper.MustBindEnv(constant.IPFSEndpoint)
	viper.MustBindEnv(constant.DIDAuthServerEndpoint)

	viper.BindEnv(constant.OperatorPrivateKey)
	viper.BindEnv(constant.OperatorPrivateKeyED25519)
}
