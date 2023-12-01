package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/coordinator/api"
	"github.com/machinefi/sprout/coordinator"
)

func main() {
	initLogger()
	bindEnvConfig()

	seq, err := coordinator.NewCoordinator(viper.GetString(DatabaseDSN), viper.GetString(P2PMultiaddr))
	if err != nil {
		log.Fatal(err)
	}
	go seq.Run()

	go func() {
		if err := api.NewHttpServer(seq).Run(viper.GetString(HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
