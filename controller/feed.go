package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/service"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	respository.Response
	VideoList []respository.Video `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

//获取视频流
func Feed(c *gin.Context) {
	//videoList, nextTime := respository.QueryByCreatedTime()
	val := c.Query("latest_time")
	start, _ := strconv.ParseInt(val, 10, 64)
	if time.Now().Unix() >= start/1000 && time.Now().Unix() <= (start/1000+500) {
		start = 0
	}
	videoList, next := service.GetVideoList(start)
	if len(videoList) == 0 {
		videoList = DemoVideos
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  respository.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  next,
	})
}
