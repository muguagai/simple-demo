package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user := respository.UsersLoginInfo[token]
	if _, exist := respository.UsersLoginInfo[token]; exist {
		videoid := c.Query("video_id")
		var isFavorite bool
		vid, _ := strconv.ParseInt(videoid, 10, 64)
		action_type := c.Query("action_type")
		var video respository.Video
		respository.Db.Where("id = ?", videoid).Find(&video)
		if action_type == "1" {
			video.FavoriteCount++
			isFavorite = true
		}
		if action_type == "2" {
			video.FavoriteCount--
			isFavorite = false
		}
		respository.Db.Save(&video)
		//查询数据库是否存在该记录
		var find bool
		userLike, find := respository.NewUserLikeDaoInstance().QueryUserLikeByVideoIDandLikeId(vid, user.Id)
		if find == true {
			userLike.IsFavorite = isFavorite
			userLike.VideoId = vid
			userLike.LikeId = user.Id
			respository.Db.Save(&userLike)
		} else {
			respository.Db.Create(&userLike)
		}
		c.JSON(http.StatusOK, respository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	user := respository.UsersLoginInfo[token]
	videos := respository.NewUserLikeDaoInstance().QueryFavoriteListByUserId(user.Id)
	if videos == nil {
		videos = DemoVideos
	}
	c.JSON(http.StatusOK, VideoListResponse{

		Response: respository.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}
