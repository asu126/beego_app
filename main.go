package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	db "app/lib/db"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.InitDb()

	// degug
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
	log.Debug("this is a debug message")

	beego.Run()
}
