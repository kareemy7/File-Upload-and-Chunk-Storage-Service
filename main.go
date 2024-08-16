package main

import (
	"File-Upload-and-Chunk-Storage-Service/controllers"
	"File-Upload-and-Chunk-Storage-Service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/upload", controllers.UploadFile)
	r.GET("/download/:file_id", controllers.DownloadFile)
	r.Run()
}
