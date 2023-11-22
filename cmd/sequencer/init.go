package main

import (
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initStdLogger() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func initEnvConfigBind() {
	viper.MustBindEnv(HttpServiceEndpoint)
	viper.MustBindEnv(GrpcServiceEndpoint)
	viper.MustBindEnv(DatabaseDSN)
}
