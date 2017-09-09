package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// m "app/models"
	// _ "github.com/go-sql-driver/mysql"
)

type HomepageController struct {
	beego.Controller
}

func (this *HomepageController) Get() {
	fmt.Println("(0000000000000 0000000000)")
	fmt.Println("(0000000000000 0000000000)")
	this.TplName = "homepage.html"
}
