package api

import (
	"fmt"
	"time"

	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Uploadfile(c *gin.Context) {
	username := c.GetString("username")
	var params requestBody
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	client := storage.Connect()

	//检查bucket是否属于用户 (?)

	// 获取上传的文件
	file, err := c.FormFile("file") // 文件字段名称为 "file"
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file: " + err.Error()})
		return
	}

	// 打开文件
	fileContent, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open file: " + err.Error()})
		return
	}
	defer fileContent.Close()

	// 使用时间戳和用户名生成唯一的文件名
	objectName := fmt.Sprintf("%s-%d-%s", username, time.Now().Unix(), file.Filename)

	// 上传文件
	fileSize := file.Size                          // 获取文件大小
	contentType := file.Header.Get("Content-Type") // 获取文件类型
	_, err = client.PutObject(c, params.BucketName, objectName, fileContent, fileSize, "", "", minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(200, gin.H{
		"message":  "File uploaded successfully",
		"fileName": objectName,
		"bucket":   params.BucketName,
	})
}
