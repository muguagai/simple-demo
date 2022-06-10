/******
** @创建时间 : 2022/6/5 11:39
** @作者 : MUGUAGAI
******/
package service

import (
	"github.com/RaymondCode/simple-demo/util"

	"github.com/RaymondCode/simple-demo/respository/redis"

	"github.com/RaymondCode/simple-demo/respository"
)

func FavoriteAction(videoid string, action_type string, user respository.User) (err error) {
	var video respository.Video
	var isFavorite bool
	respository.Db.Where("id = ?", videoid).Find(&video)
	if action_type == "1" {
		video.FavoriteCount++
		isFavorite = true
		//视频的分数增加并存储点赞记录
		redis.LikedForVideo(videoid, isFavorite, user.Id)
	}
	if action_type == "2" {
		video.FavoriteCount--
		isFavorite = false
		//视频的分数减少并删除点赞记录
		redis.LikedForVideo(videoid, isFavorite, user.Id)
	}
	respository.Db.Save(&video)
	return nil
}

func FavouriteList(userID int64) (videoList []respository.Video) {
	videoIDS := redis.GetFavouriteVideo(userID)
	vID := util.String2Int(videoIDS)
	return respository.GetVideoListByIDs(vID)
}
