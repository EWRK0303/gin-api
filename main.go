package main

import (
	"github.com/EWRK0303/gin-api/config"
	"github.com/EWRK0303/gin-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //创建路由

	database.Setup() //setup数据库

	config.Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
