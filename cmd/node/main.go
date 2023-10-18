package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/machinefi/w3bstream-mainnet/pkg/msghandler"
)

func main() {
	msgHandler := msghandler.New()

	router := gin.Default()
	router.POST("/message", func(c *gin.Context) {
		var req msgReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, newErrResp(err))
			return
		}
		msg := &msghandler.Msg{
			Data: []byte(req.Data),
		}
		if err := msgHandler.Handle(msg); err != nil {
			c.JSON(http.StatusInternalServerError, newErrResp(err))
			return
		}

		c.Status(http.StatusOK)
	})

	go func() {
		// TODO move ":9000" to config
		if err := router.Run(":9000"); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
