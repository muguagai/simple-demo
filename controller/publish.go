package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaymondCode/simple-demo/respository"

	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	respository.Response
	VideoList []respository.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := c.PostForm("token")
	if _, exist := respository.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// 获取上传文件信息
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	user := respository.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	var video = respository.Video{
		Id:      worker.GetId(),
		Author:  user,
		PlayUrl: "http://10.60.160.81:8080/static/" + finalName,
		//封面固定
		CoverUrl:      "http://10.60.160.81:8080/static/fengmian.webp",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateTime:    time.Now(),
	}
	respository.Db.Create(&video)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, respository.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user := respository.UsersLoginInfo[token]
	c.JSON(http.StatusOK, VideoListResponse{
		Response: respository.Response{
			StatusCode: 0,
		},
		VideoList: respository.QueryVideosListByauthorid(user),
	})
}
