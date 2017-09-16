package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	// Iid     int
	Name     string
	Username string
	Password string
	Profile  *Profile `orm:"rel(one)"`      // OneToOne relation
	Post     []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
	Sex      int64    `orm:"default(1)"`
	Abc      string   `orm:"null"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

// 为什么一定要大写开头?
func FindByID(id int) (User, error) {
	o := orm.NewOrm()
	user := User{Id: id, Username: "admin", Password: "admin"}

	err := o.Read(&user)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键")
	// } else {
	// 	fmt.Println(user.Id, user.Name)
	// }

	return user, err
}

func Ahtu_for_login(username, password string) (User, error) {
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")
	// 也可以直接使用对象作为表名
	// user := new(User)
	// qs = o.QueryTable(user) // 返回 QuerySeter
	fmt.Println("----------------------------")

	var user User
	err := qs.Filter("username", username).Filter("password", password).One(&user)
	return user, err
}

func InsertUser() {
	o := orm.NewOrm()

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
