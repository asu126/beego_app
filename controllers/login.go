package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	mysession "app/lib/session"
	m "app/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
	log.Debug("this is a debug message -dsds")

	keys := make([]string, 0)
	this.Ctx.Input.Bind(&keys, "keys")
	fmt.Println(keys, len(keys))
	fmt.Println(this.Ctx.Request.RequestURI)

	user, _ := m.FindByID(66)
	fmt.Println(user.Id)

	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")

	user, err := m.Ahtu_for_login(username, password)

	if err == nil {
		fmt.Println(user.Id)
		fmt.Println(user.Name)

		session_id, err := mysession.SetSessionIdByUsername(this.Ctx.ResponseWriter, this.Ctx.Request, user.Username)
		this.SetSession("BeegoSessionID", session_id)
		if err != nil {
			fmt.Println("set session error")
		}
		this.Redirect("/dashbord", 302)
	} else {
		fmt.Println(err)
		fmt.Println("login error..............")
		this.Redirect("/login", 302)
	}
}
