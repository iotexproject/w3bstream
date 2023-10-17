package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/machinefi/w3bstream-mainnet/pkg/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/pkg/receiver"
)

func main() {

	mq := gochan.New(4096)
	receiver := receiver.New(mq)

	go func() {
		if err := receiver.Run(":9000"); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
