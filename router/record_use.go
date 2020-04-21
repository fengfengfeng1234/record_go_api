package router

import (
	"github.com/gin-gonic/gin"
	"record_go_api/api"
)

func InitRecordUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/", api.GetUser)
	}
}
