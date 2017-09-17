package controllers

import (
	"github.com/astaxie/beego"
)

type HelpController struct {
	beego.Controller
}

func (this *HelpController) Get() {
	// fmt.Println("wwwwwwwwwww")
	this.Data["Website"] = this.Ctx.Input.Param(":id")
	this.Data["Email"] = "astfffffffffffffffffffffffffxie@gmail.com"
	this.TplName = "help.tpl"
}
