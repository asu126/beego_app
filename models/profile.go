package models

import (
	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Profile))
}
