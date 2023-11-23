package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/sequencer/api"
	"github.com/machinefi/sprout/sequencer"
)

func main() {
	initLogger()
	bindEnvConfig()

	seq, err := sequencer.NewSequencer(viper.GetString(DatabaseDSN))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(seq).Run(viper.GetString(HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := api.NewGrpcServer(seq).Run(viper.GetString(GrpcServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
