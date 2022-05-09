package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
/*var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Password:      "123",
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}*/

//var slice1 []int64 = make([]int64, 50)

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{

			Id:       worker.GetId(),
			Name:     username,
			Password: password,
		}
		db.Create(&newUser)
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

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
	err := db.Where("id = ?", id).Find(&user).Error
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
	err := db.Where("id in (?)", ids).Find(&users).Error
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
	result := db.Find(&users)
	if result.Error != nil {
		util.Logger.Error("find all user err:")
		return nil, result.Error
	}
	userMap := make(map[string]User)
	for _, user := range users {
		userMap[user.Name+user.Password] = user
	}
	return userMap, nil
}

func CreatUserinfo() map[string]User {
	all, _ := NewUserDaoInstance().QueryAll()
	return all
}
