package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"record_go_api/global"
	"xorm.io/xorm"
)

func Mysql() {

	if db, err := xorm.NewEngine("mysql", "root:root@(127.0.0.1:8889)/RECORD_DB?charset=utf8&parseTime=True&loc=Local"); err != nil {
		fmt.Println("DEFAULTDB数据库启动异常", err)
	} else {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(10)
		global.GVA_DB.DB().SetMaxOpenConns(10)
		global.GVA_DB.ShowSQL(true)

	}

}
