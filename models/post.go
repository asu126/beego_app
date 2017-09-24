package models

import (
	"fmt"
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

// Post 和 User 是 ManyToOne 关系
func GetUserPostList(uid int64) (posts []*Post, err error) {
	// func GetUserPostList(uid int64) {
	// var posts []*Post
	o := orm.NewOrm()
	post := new(Post)
	o.QueryTable(post).Filter("User", 1).RelatedSel().All(&posts)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Printf("%d posts read\n", num)
	for _, post := range posts {
		fmt.Printf("Id: %d, UserName: %d, Title: %s\n", post.Id, post.User.Username, post.Title)
	}
	return posts, nil
}
