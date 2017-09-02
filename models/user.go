package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	// Iid     int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
	Sex     int64    `orm:"default(1)"`
	Abc     string   `orm:"null"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

// 为什么一定要大写开头?
func FindByID(id int) (User, error) {
	o := orm.NewOrm()
	user := User{Id: id}

	err := o.Read(&user)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键")
	// } else {
	// 	fmt.Println(user.Id, user.Name)
	// }
	if err != nil {
		return user, err
	}

	fmt.Println(user.Abc)
	return user, nil
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
