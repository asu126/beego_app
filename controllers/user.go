package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	// fmt.Println("wwwwwwwwwww")
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
	log.Debug("this is a debug message -dsds")

	keys := make([]string, 0)
	this.Ctx.Input.Bind(&keys, "keys")
	fmt.Println(keys, len(keys))
	fmt.Println(this.Ctx.Request.RequestURI)

	this.Data["Website"] = this.Ctx.Input.Param(":id")
	this.Data["Email"] = "astfffffffffffffffffffffffffxie@gmail.com"
	this.TplName = "index.tpl"
}
