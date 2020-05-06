package initialize

import (
	"github.com/gin-gonic/gin"
	"record_go_api/router"
)

//初始化总路由
func routers() *gin.Engine {
	var Router = gin.Default()
	ApiGroup := Router.Group("")
	//use api
	router.InitRecordUserRouter(ApiGroup)
	router.InitRecordTimeLineRouter(ApiGroup)
	return Router
}

func RunWindowsServer() {
	var router = routers()
	router.Run(":8080")
}
