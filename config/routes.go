package config

import (
	"github.com/EWRK0303/gin-api/api"
	"github.com/EWRK0303/gin-api/midware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	//数据库操作
	r.GET("/index", midware.Auth(), api.Index) //客户端请求‘index’时，返回json格式的全部查询结果
	r.POST("/create", api.Create)              //客户端post ‘create’时，返回添加数据行 （404报错）
	r.POST("/delete/:id", api.Delete)

	//用户登陆
	r.POST("/login", api.Login)

	//bucket操作
	r.POST("/createbucket", midware.Auth(), api.Createbucket) //用户创建bucket
	r.POST("/deletebucket", midware.Auth(), api.Deletebucket) //删除bucket
	r.GET("/showbucket", midware.Auth(), api.ShowBucket)      //显示用户所有bucket

}
