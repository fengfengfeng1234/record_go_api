package api

import (
	"github.com/gin-gonic/gin"
	"record_go_api/global/response"
	"record_go_api/model"
)

/**
  查询数据列表
*/
func QueryTimeLineList(c *gin.Context) {

	var req *model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("解析出错", c)
		return
	}

	response.OkDetailed("timeLine", "注册成功", c)
}
