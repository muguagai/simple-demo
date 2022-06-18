package respository

import (
	"errors"
	"sync"

	"github.com/RaymondCode/simple-demo/util"
	"gorm.io/gorm"
)

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := Db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
	var users []*User
	err := Db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		util.Logger.Error("batch find user by id err:" + err.Error())
		return nil, err
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}

func (*UserDao) QueryAll() (map[string]User, error) {
	var users []User
	result := Db.Find(&users)
	if result.Error != nil {
		util.Logger.Error("find all user err:")
		return nil, result.Error
	}
	userMap := make(map[string]User)
	for _, user := range users {
		userMap[user.Name] = user
	}
	return userMap, nil
}

func (*UserDao) SaveUser(user User) error {
	err := Db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (*UserDao) CheckUserExist(username string) (error error) {
	var user User
	row := Db.Where("name = ?", username).First(&user).RowsAffected
	if row > 0 {
		return errors.New("用户已存在")
	}
	return
}

func (*UserDao) QueryUserByUserName(username string) (user User) {
	row := Db.Where("name = ?", username).First(&user).RowsAffected
	if row == 0 {
		return
	}
	return user
}
