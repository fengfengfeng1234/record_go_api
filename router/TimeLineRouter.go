package router

import (
	"github.com/gin-gonic/gin"
	"record_go_api/api"
)

func InitRecordTimeLineRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("timeLine")
	{
		UserRouter.GET("/", api.QueryTimeLineList)
	}
}
