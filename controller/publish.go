package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var token string

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//没有token参数，固定一个token
	token := "feng123456"
	//token = t.String()
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// 获取上传文件信息
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	var video = Video{
		Id:            worker.GetId(),
		Author:        user,
		PlayUrl:       "http://192.168.1.5:8080/static/" + finalName,
		CoverUrl:      "http://192.168.1.5:8080/static/屏幕截图 2021-02-16 163146.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateTime:    time.Now(),
	}
	db.Create(&video)
	//videoInfo[video.Id]:=video
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token = c.Query("token")
	user := usersLoginInfo[token]
	cookie, _ := c.Cookie("name")
	fmt.Println(cookie)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: QueryList(user),
	})
}

func QueryList(user User) []Video {
	var Videos []Video
	db.Where("author_id in (?)", user.Id).Find(&Videos)
	return Videos
}
