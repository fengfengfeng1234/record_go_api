package api

import (
	"github.com/gin-gonic/gin"
	"record_go_api/global/response"
)

func GetUser(c *gin.Context)  {
 	 response.OkDetailed("获取用户信息", "注册成功", c)
}
