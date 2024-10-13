package api

import (
	"github.com/EWRK0303/gin-api/database"
	"github.com/EWRK0303/gin-api/models"
	"github.com/gin-gonic/gin"
)

type Createparams struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func Create(c *gin.Context) {
	var params Createparams //传入一个createparams结构体，定义如上
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: params.Name, Price: params.Price}
	database.GetDB().Create(&product)

	c.JSON(200, gin.H{"product": product})

}
