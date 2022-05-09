package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var usersLoginInfo map[string]User

func Init() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	usersLoginInfo = CreatUserinfo()
	return err
}
