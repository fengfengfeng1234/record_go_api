package main

import (
	"fmt"
	"record_go_api/initialize"
	"record_go_api/model/recordTable"
)

//  https://tour.go-zh.org/moretypes/1  指针
func main() {
	// 初始化数据库
	initialize.Mysql()
	//initialize.RunWindowsServer()

	var info recordTable.Token
	//result, error := info.One()
	//result, error := info.CreateTokenModel(1121)
	//result, error := info.GetToken("5f4deeffd43e4e38ae214d2869e9fe9b")
	//result, error := info.CheckToken("5f4deeffd43e4e38ae214d2869e9fe9b")
	result, error := info.DeleteToken(1121)
	fmt.Println(result)
	fmt.Println(error)

}
