package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Uploadfile(c *gin.Context) {
	// username := c.GetString("username")

	bucketName := c.PostForm("bucket_name")
	objectName := c.PostForm("object_name")

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

	// 上传文件
	fileSize := file.Size                          // 获取文件大小
	contentType := file.Header.Get("Content-Type") // 获取文件类型
	_, err = client.PutObject(c, bucketName, objectName, fileContent, fileSize, "", "", minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	//添加到数据库
	file_name := models.File{BucketId: bucketName, FileName: objectName}
	database.GetDB().Create(&file_name)

	// 返回成功响应
	c.JSON(200, gin.H{
		"message":  "File uploaded successfully",
		"fileName": objectName,
		"bucket":   bucketName,
	})
}
