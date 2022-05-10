package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user := usersLoginInfo[token]
	if _, exist := usersLoginInfo[token]; exist {
		videoid := c.Query("video_id")
		var isFavorite bool
		vid, _ := strconv.ParseInt(videoid, 10, 64)
		action_type := c.Query("action_type")
		var video Video
		db.Where("id = ?", videoid).Find(&video)
		if action_type == "1" {
			video.FavoriteCount++
			isFavorite = true
		}
		if action_type == "2" {
			video.FavoriteCount--
			isFavorite = false
		}
		db.Save(&video)
		//查询数据库是否存在该记录
		var userLike UserLike
		find := db.Table("user_likes").Where("video_id = ?", videoid).Where("like_id = ?", user.Id).Find(&userLike)
		if find != nil {
			userLike.IsFavorite = isFavorite
			userLike.VideoId = vid
			userLike.LikeId = user.Id
			db.Save(&userLike)
		} else {
			db.Create(&userLike)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	user := usersLoginInfo[token]
	videos := QueryFavoriteListByUserId(user.Id)
	if videos == nil {
		videos = DemoVideos
	}
	c.JSON(http.StatusOK, VideoListResponse{

		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}

//根据userid查询用户喜爱列表
func QueryFavoriteListByUserId(userId int64) []Video {
	var Videos []Video = nil
	var videoIds []*int64
	db.Table("user_likes").Select("video_id").Where("like_id = ?", userId).Where("is_favorite= ?", 1).Find(&videoIds)
	if len(videoIds) == 0 {
		return nil
	}
	db.Find(&Videos, videoIds)
	return Videos
}
