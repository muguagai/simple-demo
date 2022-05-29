package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/respository"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	respository.Response
	VideoList []respository.Video `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

//获取视频流
// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videoList, nextTime := respository.QueryByCreatedTime()
	if len(videoList) == 0 {
		videoList = DemoVideos
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  respository.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
