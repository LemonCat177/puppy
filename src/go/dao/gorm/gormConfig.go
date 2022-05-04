package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/*
author Lei Wang
本文件中需导入(import)数据库驱动包
*/

/*
为了便于说明,conn和dialect常量定义在此文件中,在项目中可将这两个常量与其他常量一起统一定义在config目录下
conn 数据库连接语句 protocol为协议类型
常用参数(parameter) value表示参数的值
charset,编码类型,应与对应的数据库设置的编码类型保持一致,一般为utf8mb4
parseTime,是否对时间格式化,一般为true
loc,对数据库进行操作时,存入数据库的时间可能会和本地时间相差8个小时,此时可用此参数进行设置,一般为Local
*/
const conn = "username:password@protocol(IpAddress:Port)/NameOfDatabase?parameter1=value&parameter2=value"

// dialect 数据库类型，例如MySQL、Oracle等，本示例中将其设置为空字符串
const dialect = ""

// InitDB 获取并返回数据库连接
func InitDB() *gorm.DB {
	db, err := gorm.Open(dialect, conn)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
