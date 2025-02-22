// 网关路由初始化与注册
// @author MoGuQAQ
// @version 1.0.0

package router

import (
	"gin-template/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	GinMode := "release"
	gin.SetMode(GinMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	Register(router)
	return router
}

func Register(router *gin.Engine) {
	router.GET("/", api.HelloWorld)
}
