package respository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var UsersLoginInfo map[string]User

//初始化数据库
func Init() error {
	var err error
	dsn := "root:123@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=false"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	var user User
	UsersLoginInfo = make(map[string]User)
	UsersLoginInfo["0"] = user
	return err
}

/*func CreatUserinfo() map[string]User {
	all, _ := NewUserDaoInstance().QueryAll()
	return all
}*/
