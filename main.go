package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	go hub.run()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/upload", func(c *gin.Context) {
		_, header, _ := c.Request.FormFile("file")
		fileName := generateUniqueFileName(header.Filename)
		c.SaveUploadedFile(header, "/data/images/"+fileName)
		c.JSON(http.StatusOK, gin.H{
			"fileName": fileName,
		})
	})
	r.GET("/ws", func(context *gin.Context) {
		serverWs(hub, context.Writer, context.Request)
	})
	err := r.Run()
	if err != nil {
		return
	}

}

func generateUniqueFileName(headerFilename string) string {
	randomComponent := rand.Intn(1000)
	fileName := fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), randomComponent, headerFilename)
	return fileName
}
