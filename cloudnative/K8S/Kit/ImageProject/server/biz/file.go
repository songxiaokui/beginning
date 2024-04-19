package biz

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

const fileNameField = "file"

func UploadHandle(c *gin.Context) {
	file, err := c.FormFile(fileNameField)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savePath := filepath.Join("./static", file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "path": savePath})
}
