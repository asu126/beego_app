package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// m "app/models"
	// _ "github.com/go-sql-driver/mysql"
	mysession "app/lib/session"
)

type DashbordController struct {
	beego.Controller
}

func (this *DashbordController) Prepare() {
	fmt.Println("Prepare func exec ...")
	if client_sid, ok := this.GetSession("BeegoSessionID").(string); ok {
		/* act on str */
		fmt.Println(client_sid)

		username, err := mysession.GetSessionIdBySid(client_sid)
		if err != nil || client_sid == "" {
			this.Redirect("/login", 302)
		} else {

			this.Data["Username"] = username
		}
	} else {
		/* not string */
		this.Redirect("/login", 302)
	}
}

func (this *DashbordController) Get() {

	this.TplName = "dashbord.html"

}

func (this *DashbordController) Reports() {
	// fmt.Println("wwwwwwwwwww")
	this.Data["Website"] = this.Ctx.Input.Param(":id")
	this.Data["Email"] = "astfffffffffffffffffffffffffxie@gmail.com"
	this.TplName = "help.tpl"
}
