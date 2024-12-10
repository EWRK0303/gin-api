package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBucket(c *gin.Context) {

	var buckets []models.Bucket
	result := database.GetDB().Where("userid = ?", c.GetString("username")).Find(&buckets)

	// 错误检查
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// 返回查询到的结果
	c.JSON(200, gin.H{"buckets": buckets})

}
