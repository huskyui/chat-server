package main

import (
	"chat-server/common"
	"chat-server/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	go websocket.WebSocketHub.Run()

	r.GET("/ping", common.Ping)
	r.GET("/upload", common.UploadFile)
	r.GET("/ws", websocket.Websocket)
	err := r.Run()
	if err != nil {
		return
	}

}
