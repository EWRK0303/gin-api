package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Deletebucket(c *gin.Context) {
	client := storage.Connect()
	var params bucketname
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := client.RemoveBucket(context.Background(), params.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"bucket deleted": params.Name})

}
