package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Deletebucket(c *gin.Context) {
	username := c.GetString("username")

	var params requestBody
	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// 检查 Bucket 是否属于当前用户 (doesn't work ? 无法触发这个403)
	if !IsBucketOwnedByUser(username, params.BucketName) {
		c.JSON(403, gin.H{"error": "You don't have permission to delete this bucket"})
		return
	}

	client := storage.Connect()

	err := client.RemoveBucket(context.Background(), params.BucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//删除数据库中的记录(userid)
	database.GetDB().Delete(&models.Bucket{}, "Userid = ? AND Bucketname = ?", username, params.BucketName)
	c.JSON(http.StatusOK, gin.H{"bucket deleted": params.BucketName})

}

// 检查 Bucket 是否属于用户
func IsBucketOwnedByUser(username, bucketName string) bool {
	var count int64
	database.GetDB().Model(&models.Bucket{}).
		Where("username = ? AND bucket_name = ?", username, bucketName).
		Count(&count)
	return count > 0
}
