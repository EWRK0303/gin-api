package main

import (
	"github.com/EWRK0303/gin-api/config"
	"github.com/EWRK0303/gin-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //创建路由

	database.Setup() //setup数据库

	config.Routes(r) //(包含'ping', 'index':查询数据库, 'create'：添加数据行进入数据库)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
