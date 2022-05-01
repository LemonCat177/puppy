package gorm

import (
	"../../model"
	"fmt"
	"github.com/jinzhu/gorm"
)

/*
author Lei Wang
*/

func InsertExample(username string, password string) {
	db := InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}(db)
	u := model.User{
		Username: username,
		Password: password,
	}
	//AutoMigrate 若结构体无对应数据表，则会自动创建数据表
	db.AutoMigrate(&u)
	//Create 插入一个数据
	db.Create(&u)
}
