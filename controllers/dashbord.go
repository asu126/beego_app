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

func (this *DashbordController) Get() {

	if client_sid, ok := this.GetSession("BeegoSessionID").(string); ok {
		/* act on str */
		fmt.Println(client_sid)

		id, err := mysession.GetSessionIdBySid(client_sid)
		if err == nil && id != "" {
			fmt.Println(id)
			this.TplName = "dashbord.html"
		} else {
			this.Redirect("/login", 302)
		}
	} else {
		/* not string */
		this.Redirect("/login", 302)
	}
}
