package api

import (
	"fmt"
	"net/http"

	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/gin-gonic/gin"
)

func ShowBucket(c *gin.Context) {
	client := storage.Connect()
	a, _ := client.ListBuckets(c)
	fmt.Println(client.ListBuckets(c))
	c.JSON(http.StatusOK, gin.H{"data": a})
}
