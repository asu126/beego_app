package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// m "app/models"
	// _ "github.com/go-sql-driver/mysql"
)

type VueController struct {
	beego.Controller
}

func (this *VueController) Get() {
	fmt.Println("(0000000000000 0000000000)")
	fmt.Println("(0000000000000 0000000000)")
	this.TplName = "vue.html"
}
