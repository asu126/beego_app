package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// m "app/models"
	// _ "github.com/go-sql-driver/mysql"
)

type DashbordController struct {
	beego.Controller
}

func (this *DashbordController) Get() {
	fmt.Println("(0000000000000 0000000000)")
	fmt.Println(this.Input().Get("username"))
	fmt.Println(this.Input().Get("password"))
	fmt.Println("(0000000000000 0000000000)")
	this.TplName = "dashbord.html"
}
