package controllers

import (
	"github.com/astaxie/beego"

	mysession "app/lib/session"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	mysession.DeleteSessionIdBySid(this.Ctx.ResponseWriter, this.Ctx.Request)
	this.Redirect("/login", 302)
}
