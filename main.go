package main

import (
	"record_go_api/initialize"

	//_ "github.com/gorilla/mux"
)

func main() {
	// 初始化数据库
	initialize.Mysql()
	initialize.RunWindowsServer()

}
