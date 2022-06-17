package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/service"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	respository.Response
	UserList []respository.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
//关注功能
func RelationAction(c *gin.Context) {
	to_user_id := c.Query("to_user_id")
	touserid, _ := strconv.ParseInt(to_user_id, 10, 64)
	token := c.Query("token")
	follower := respository.UsersLoginInfo[token]
	actiontype := c.Query("action_type")
	var follow_follower respository.FollowFollower
	//作者是否存在
	follow, err := respository.NewUserDaoInstance().QueryUserById(touserid)
	find := respository.Db.Table("follow_followers").Where("follow_id = ?", follow.Id).Where("follower_id = ?", follower.Id).Find(&follow_follower)
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "Author doesn't exist"})
	}
	if touserid == follower.Id {
		//作者不能关注自己
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User is Author"})
	} else if _, exist := respository.UsersLoginInfo[token]; exist {
		//用户存在
		service.RelationAction(follow, actiontype, follower, follow_follower, find)
	} else {
		//用户不存在
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userid := c.Query("user_id")
	uid, _ := strconv.ParseInt(userid, 10, 64)
	FollowList := respository.QueryFollowListByUserId(uid)
	c.JSON(http.StatusOK, UserListResponse{
		Response: respository.Response{
			StatusCode: 0,
		},
		UserList: FollowList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userid := c.Query("user_id")
	uid, _ := strconv.ParseInt(userid, 10, 64)
	FollowerList := respository.QueryFollowerListByUserId(uid)
	c.JSON(http.StatusOK, UserListResponse{
		Response: respository.Response{
			StatusCode: 0,
		},
		UserList: FollowerList,
	})
}
