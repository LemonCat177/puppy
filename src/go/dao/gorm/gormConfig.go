package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
author Lei Wang
*/

// conn 数据库连接语句 protocol为协议类型
const conn = "username:password@protocol(IpAddress:Port)/NameOfDatabase?parameter1&parameter2"

// dialect 数据库类型 例如MySQL、Oracle等
const dialect = ""

// InitDB 获取并返回数据库连接
func InitDB() *gorm.DB {
	db, err := gorm.Open(dialect, conn)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return db
}
