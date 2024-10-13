package config

import (
	"github.com/EWRK0303/gin-api/api"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", api.Ping)      //客户端请求‘ping’时，返回json格式的‘pong44’
	r.GET("/index", api.Index)    //客户端请求‘index’时，返回json格式的全部查询结果
	r.POST("/create", api.Create) //客户端post ‘create’时，返回
}
