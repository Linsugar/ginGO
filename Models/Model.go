package Models

import "github.com/jinzhu/gorm"

//登录
type Login struct {
	/*
			Card身份证，Name昵称，Address地址，Age年龄，Ip地址，
		UserName用户名，UserPwd用户密码，Photo头像,Statue是否登录
	*/
	gorm.Model
	Card     string
	Name     string
	Address  string
	Age      int
	Ip       string
	UserName string
	UserPwd  string
	Photo    string
	Statue   bool
}

//文章
type WebTitle struct {
	/*
			Title标题，Author作者,Text文本，
		IsDelete是否删除,TextImage文本封面
	*/
	gorm.Model
	Title     string
	Author    string
	Text      string
	IsDelete  bool
	TextImage string
}

//日志记录
type History struct {
	/*
			User谁，Advice建议,IsDelete是否删除，
		IsDelete是否删除,Title标题，Image头像
	*/
	gorm.Model
	User     string
	Advice   string
	IsDelete bool
	Title    string
	Image    string
}
