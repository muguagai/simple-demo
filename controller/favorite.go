package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/respository/redis"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	user := respository.UsersLoginInfo[token]
	if _, exist := respository.UsersLoginInfo[token]; exist {
		videoid := c.Query("video_id")
		action_type := c.Query("action_type")
		service.FavoriteAction(videoid, action_type, user)
		c.JSON(http.StatusOK, respository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	user := respository.UsersLoginInfo[token]
	//videos := respository.NewUserLikeDaoInstance().QueryFavoriteListByUserId(user.Id)
	videos := service.FavouriteList(user.Id)
	if videos == nil {
		videos = DemoVideos
	}
	redis.GetFavouriteVideo(user.Id)
	c.JSON(http.StatusOK, VideoListResponse{

		Response: respository.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}
