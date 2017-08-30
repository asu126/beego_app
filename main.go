package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
	log.Debug("this is a debug message")
	beego.Run()
}
