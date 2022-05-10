package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: QueryByCreatedTime(),
		NextTime:  time.Now().Unix(),
	})
}

func QueryByCreatedTime() []Video {
	var Videos []Video
	db.Table("videos").Order("create_time desc").Limit(30).Find(&Videos)
	//将Author和Video批量关联
	//不能使用range遍历
	for i := 0; i < len(Videos); i++ {
		var user User
		db.Where("id = ?", Videos[i].AuthorID).Find(&user)
		Videos[i].Author = user
	}
	return Videos
}

/*func CreatVideoinfo() map[string]User {
	videos := QueryByCreatedTime()
}*/
