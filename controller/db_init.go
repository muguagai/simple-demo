package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var usersLoginInfo map[string]User
var videoInfo map[string]Video

func Init() error {
	var err error
	dsn := "root:123@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	usersLoginInfo = CreatUserinfo()
	return err
}
