package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/cmd/node/apis/message"
	"github.com/machinefi/w3bstream-mainnet/enums"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var engine = gin.Default()

func init() {
	message.Register(engine)
}

func Run() {
	go func() {
		if err := engine.Run(viper.GetString(enums.EnvKeyServiceEndpoint)); err != nil {
			slog.Error("serving http: ", err)
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
