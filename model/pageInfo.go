package model

//分页请求参数
type SelectPageReq struct {
	PageNum  int `form:"pageNum"`  //当前页码
	PageSize int `form:"pageSize"` //每页数
}
