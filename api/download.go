package api

import (
	"fmt"
	"io"

	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func DownloadFile(c *gin.Context) {
	// username := c.GetString("username")

	bucketName := c.PostForm("bucket_name")
	objectName := c.PostForm("object_name")

	client := storage.Connect()

	object, ObjectInfo, _, err := client.GetObject(c, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", objectName))
	c.Header("Content-Length", fmt.Sprintf("%d", ObjectInfo.Size))
	c.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, object)
		if err != nil {
			c.Error(err)
			return false
		}
		return true
	})

}
