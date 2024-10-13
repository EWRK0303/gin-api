package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var id = c.Param("id")
	database.GetDB().Delete(&models.Product{}, id)
	c.JSON(200, gin.H{"message": "deleted"})
}
