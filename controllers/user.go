package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	m "app/models"
	// _ "github.com/go-sql-driver/mysql"
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

	users, count := m.Getuserlist(0, 10, "Id")
	posts, _ := m.GetUserPostList(1)
	this.Data["users"] = &users
	this.Data["posts"] = &posts
	fmt.Println(users)
	fmt.Println(count)

	this.Data["Website"] = this.Ctx.Input.Param(":id")
	this.Data["Email"] = "astfffffffffffffffffffffffffxie@gmail.com"
	this.TplName = "users.html"
}
