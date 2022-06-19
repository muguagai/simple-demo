package controller

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/RaymondCode/simple-demo/util/jwt"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
/*var UsersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Password:      "123",
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}*/

var userIdSequence = int64(1)
var userDao respository.UserDao

type UserLoginResponse struct {
	respository.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	respository.Response
	User respository.User `json:"user"`
}

func Register(c *gin.Context) {
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	username := c.Query("username")
	password := c.Query("password")
	if userDao.CheckUserExist(username) != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: respository.Response{StatusCode: 1, StatusMsg: "用户名已存在"},
		})
	} else {
		//自增ID
		atomic.AddInt64(&userIdSequence, 1)
		newUser := respository.User{
			Id:       worker.GetId(),
			Name:     username,
			Password: util.MD5(password),
		}
		respository.UsersLoginInfo[newUser.Name] = newUser
		respository.Db.Create(&newUser)
		token, err := jwt.GenToken(username, password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: respository.Response{StatusCode: 1, StatusMsg: "生成token失败"},
			})
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: respository.Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	password = util.MD5(password)
	user := respository.NewUserDaoInstance().QueryUserByUserName(username)
	if len(user.Name) != 0 {
		if user.Password == password {
			token, err := jwt.GenToken(username, password)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: respository.Response{StatusCode: 1, StatusMsg: "生成token失败"},
				})
			}
			respository.UsersLoginInfo[user.Name] = user
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: respository.Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: respository.Response{StatusCode: 1, StatusMsg: "密码错误"},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	parseToken, _ := jwt.ParseToken(token)
	username := parseToken.Username
	if user, exist := respository.UsersLoginInfo[username]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: respository.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
