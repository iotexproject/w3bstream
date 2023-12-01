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
	viper.MustBindEnv(HttpServiceEndpoint)
	viper.MustBindEnv(GrpcServiceEndpoint)
	viper.MustBindEnv(P2PMultiaddr)
	viper.MustBindEnv(DatabaseDSN)
}
