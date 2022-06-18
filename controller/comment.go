package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/util/jwt"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	respository.Response
	CommentList []respository.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	parseToken, _ := jwt.ParseToken(token)
	username := parseToken.Username
	userid := c.PostForm("user_id")
	userid = c.Query("user_id")
	fmt.Println(userid)
	text := c.Query("comment_text")
	user := respository.UsersLoginInfo[username]
	if _, exist := respository.UsersLoginInfo[username]; exist {
		videoid := c.Query("video_id")
		vid, _ := strconv.ParseInt(videoid, 10, 64)
		action_type := c.Query("action_type")
		var video respository.Video
		respository.Db.Where("id = ?", videoid).Find(&video)
		if action_type == "1" {
			video.CommentCount++
			respository.Db.Save(&video)
			//添加评论
			comment := respository.Comment{
				UserID:     user.Id,
				User:       user,
				VideoID:    vid,
				Video:      video,
				Content:    text,
				CreateDate: time.Now().String(),
			}
			respository.Db.Save(&comment)

		}
		if action_type == "2" {
			//删除评论
		}

		c.JSON(http.StatusOK, respository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoid := c.Query("video_id")
	vid, _ := strconv.ParseInt(videoid, 10, 64)
	comments := respository.QueryCommentListByVideoid(vid)
	for i := 0; i < len(comments); i++ {
		var user respository.User
		respository.Db.Where("id = ?", comments[i].UserID).Find(&user)
		comments[i].User = user
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    respository.Response{StatusCode: 0},
		CommentList: comments,
	})
}
