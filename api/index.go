package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//var product []models.Product //定义变量
	var bucket []models.Bucket
	var file []models.File
	//database.GetDB().Find(&product) //检索全部对象
	database.GetDB().Find(&bucket)
	database.GetDB().Find(&file)
	//c.JSON(200, gin.H{"list": product}) //返回JSON 格式的list:product
	c.JSON(200, gin.H{"list": bucket})
	c.JSON(200, gin.H{"list": file})
}
