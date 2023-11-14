package main

import (
	msghandler "github.com/machinefi/w3bstream-mainnet/msg/handler"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/machinefi/w3bstream-mainnet/msg"
)

func main() {
	router := gin.Default()
	router.POST("/message", func(c *gin.Context) {
		var req msgReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, newErrResp(err))
			return
		}
		msg := &msg.Msg{
			ProjectID:      req.ProjectID,
			ProjectVersion: req.ProjectVersion,
			Data:           req.Data,
		}
		slog.Debug("received your message, handling")
		if err := msghandler.DefaultHandler.Handle(msg); err != nil {
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
