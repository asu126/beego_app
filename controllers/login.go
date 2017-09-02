package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	m "app/models"
	// _ "github.com/go-sql-driver/mysql"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// fmt.Println("wwwwwwwwwww")
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
	fmt.Println("(0000000000000 0000000000)")
	fmt.Println(this.Input().Get("username"))
	fmt.Println(this.Input().Get("password"))
	fmt.Println("(0000000000000 0000000000)")
	this.TplName = "dashbord.html"
}
