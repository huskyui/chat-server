package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	go hub.run()

	r.GET("/ws", func(context *gin.Context) {
		serverWs(hub, context.Writer, context.Request)
	})
	err := r.Run()
	if err != nil {
		return
	}

}
