package gorm

import (
	"../../model"
)

/*
author Lei Wang
*/

func Register(username string, password string) uint {
	db := InitDB()
	//本条语句可用下面被注释掉的匿名函数代替
	defer CloseDB(db)
	/*
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}(db)
	*/
	u := model.User{
		Username: username,
		Password: password,
	}
	//AutoMigrate 若结构体无对应数据表，则会自动创建数据表
	db.AutoMigrate(&u)
	//Create 插入一个数据
	db.Create(&u)
	return u.ID
}
