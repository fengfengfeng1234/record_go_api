package main

import (
	"fmt"
	"record_go_api/initialize"
	"record_go_api/router/middleware/anthorization"
)

func main() {
	// 初始化数据库
	initialize.Mysql()
	//initialize.RunWindowsServer()

	if result, err := anthorization.GetToken("720ad1c464cd44458f4d0649885b37f8"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.UserId)
		fmt.Println(result.Token)
		fmt.Println(result.CreateTime)
		fmt.Println(result.InvalidTime)

	}

}
