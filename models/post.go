package models

import (
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Post))
}
