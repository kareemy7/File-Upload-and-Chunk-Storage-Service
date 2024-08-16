package controllers

import (
	"File-Upload-and-Chunk-Storage-Service/initializers"
	"File-Upload-and-Chunk-Storage-Service/models"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file_")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer f.Close()

	chunkSize := 1024 * 1024
	buffer := make([]byte, chunkSize)
	numChunks := 0

	fileID := uuid.New().String()

	chunkIDs := make([]string, 0)

	for {
		n, err := f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			break
		}
		if n == 0 {
			break
		}

		chunkID := uuid.New().String()

		err = initializers.DB.Put([]byte(chunkID), buffer[:n], nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		chunkIDs = append(chunkIDs, chunkID)

		numChunks++
	}
	fileMetadata := models.FileMetadata{
		FileID:     fileID,
		FileName:   file.Filename,
		ChunkSize:  chunkSize,
		NumChunks:  numChunks,
		UploadDate: time.Now(),
		ChunkIDs:   chunkIDs,
	}
	fileMetadataBytes, err := json.Marshal(fileMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = initializers.DB.Put([]byte("file_metadata_"+fileID), fileMetadataBytes, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file_id": fileID,
	})
}

func DownloadFile(c *gin.Context) {
	fileID := c.Param("file_id")
	fileMetadataKey := []byte("file_metadata_" + fileID)
	fileMetadataValue, err := initializers.DB.Get(fileMetadataKey, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fileMetadata := models.FileMetadata{}
	err = json.Unmarshal(fileMetadataValue, &fileMetadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var fileData []byte

	for _, chunkID := range fileMetadata.ChunkIDs {
		chunkKey := []byte(chunkID)
		chunkValue, err := initializers.DB.Get(chunkKey, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		fileData = append(fileData, chunkValue...)
	}
	contentType := mime.TypeByExtension(filepath.Ext(fileMetadata.FileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	c.Header("Content-Disposition", "attachment; filename=\""+fileMetadata.FileName+"\"")
	c.Data(http.StatusOK, contentType, fileData)

}
