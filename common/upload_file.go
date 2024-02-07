package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func UploadFile(c *gin.Context) {
	_, header, _ := c.Request.FormFile("file")
	fileName := generateUniqueFileName(header.Filename)
	c.SaveUploadedFile(header, "/data/images/"+fileName)
	c.JSON(http.StatusOK, gin.H{
		"fileName": fileName,
	})
}

func generateUniqueFileName(headerFilename string) string {
	randomComponent := rand.Intn(1000)
	fileName := fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), randomComponent, headerFilename)
	return fileName
}
