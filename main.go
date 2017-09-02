package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	// 我的mysql的root用户密码为tom，打算把数据表建立在名为test数据库里
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8", 30)

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
