package model

import "github.com/jinzhu/gorm"

/*
author Lei Wang
*/

// User 自定义结构体，引用了gorm框架的Model
type User struct {
	gorm.Model
	// form 当前端提交的数据为form-data格式时对应的http请求变量名
	// json 当前端提交的数据为JSON格式时对应的http请求变量名
	// binding 当前端提交数据时此变量是否为必填
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
