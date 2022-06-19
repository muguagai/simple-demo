package controller

import (
	"bytes"
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/util/jwt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/RaymondCode/simple-demo/service"

	"github.com/RaymondCode/simple-demo/respository"

	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VideoListResponse struct {
	respository.Response
	VideoList []respository.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	worker, err := util.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := c.PostForm("token")
	parseToken, err := jwt.ParseToken(token)
	username := parseToken.Username
	if _, exist := respository.UsersLoginInfo[username]; !exist {
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
	user := respository.UsersLoginInfo[username]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	guid := uuid.New()
	guidStr := guid.String()
	coverFile := filepath.Join("./public/", guidStr+"_cover.png")

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, respository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 抽取视频封面 ffmpeg - start
	exec_ffmpeg_extract_cmd := "ffmpeg -i " + saveFile + " -ss 00:00:00 -frames:v 1 " + coverFile
	println("to tun:", exec_ffmpeg_extract_cmd)

	cmdArguments := []string{"-i", saveFile, "-ss", "00:00:00",
		"-frames:v", "1", coverFile}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	errFFMPEG := cmd.Run()
	if errFFMPEG != nil {
		log.Fatal(errFFMPEG)
	}
	fmt.Printf("command output: %q", out.String())
	// 抽取视频封面 ffmpeg - end
	var video = respository.Video{
		Id:            worker.GetId(),
		Author:        user,
		PlayUrl:       "http://" + config.Ip.String() + ":8080/static/" + finalName,
		CoverUrl:      "http://" + config.Ip.String() + ":8080/static/" + guidStr + "_cover.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateTime:    time.Now(),
		Title:         title,
	}

	if err := service.PublishVideo(video); err != nil {
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
	parseToken, _ := jwt.ParseToken(token)
	username := parseToken.Username
	user := respository.UsersLoginInfo[username]
	c.JSON(http.StatusOK, VideoListResponse{
		Response: respository.Response{
			StatusCode: 0,
		},
		VideoList: respository.QueryVideosListByauthorid(user),
	})
}
