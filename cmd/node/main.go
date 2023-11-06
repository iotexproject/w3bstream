package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/msg/handler"
	"github.com/machinefi/w3bstream-mainnet/vm"
)

func main() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	dbMigrate()
	viper.MustBindEnv("ENDPOINT")
	viper.MustBindEnv("RISC0_SERVER_ENDPOINT")
	viper.MustBindEnv("PROJECT_CONFIG_FILE")
	viper.MustBindEnv("CHAIN_ENDPOINT")
	viper.MustBindEnv("OPERATOR_PRIVATE_KEY")

	vmHandler := vm.NewHandler(viper.Get("RISC0_SERVER_ENDPOINT").(string), viper.Get("PROJECT_CONFIG_FILE").(string))
	msgHandler := handler.New(vmHandler, viper.Get("CHAIN_ENDPOINT").(string), viper.Get("OPERATOR_PRIVATE_KEY").(string))

	router := gin.Default()
	router.POST("/message", func(c *gin.Context) {
		var req msgReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, newErrResp(err))
			return
		}
		msg := &msg.Msg{
			Data: []byte(req.Data),
		}
		slog.Debug("received your message, handling")
		if err := msgHandler.Handle(msg); err != nil {
			c.JSON(http.StatusInternalServerError, newErrResp(err))
			return
		}

		c.Status(http.StatusOK)
	})

	go func() {
		if err := router.Run(viper.Get("ENDPOINT").(string)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
