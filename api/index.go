package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var product []models.Product        //定义变量
	database.GetDB().Find(&product)     //检索全部对象
	c.JSON(200, gin.H{"list": product}) //返回JSON 格式的list:product
}
