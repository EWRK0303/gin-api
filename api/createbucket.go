package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EWRK0303/gin-api/pkg/storage"
	"github.com/minio/minio-go/v7"

	"github.com/gin-gonic/gin"
)

type bucketname struct {
	Name string `json:"name"`
}

func Createbucket(c *gin.Context) {
	client := storage.Connect()
	var params bucketname
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//创建一个bucket
	err := client.MakeBucket(context.Background(), params.Name, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	//添加到数据库(?)
	fmt.Println("Bucket successfully created.")
	c.JSON(http.StatusOK, gin.H{"bucket created": params.Name})

}
