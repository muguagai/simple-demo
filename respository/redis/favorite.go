/******
** @创建时间 : 2022/6/17 21:04
** @作者 : MUGUAGAI
******/
package redis

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/respository"
	"strconv"
)

//定时将redis中的缓存存入mysql中
func FavouriteToMysql() {
	fmt.Println("清除redis数据")
	//循环获取redis中用户点赞信息
	for {
		var video respository.Video
		if client.SCard("douyin:user:liked:").Val() == 0 {
			break
		}
		videoID := client.SPop("douyin:user:liked:").Val()

		respository.Db.Where("id = ?", videoID).Find(&video)
		//获取redis中视频点赞信息保存进数据库
		for {
			if client.SCard("douyin:video:liked:"+videoID).Val() == 0 {
				break
			}
			var userlike respository.UserLike
			userID := client.SPop("douyin:video:liked:" + videoID).Val()

			userlike.VideoId, _ = strconv.ParseInt(videoID, 10, 64)
			userlike.LikeId, _ = strconv.ParseInt(userID, 10, 64)
			userlike.IsFavorite = true
			video.FavoriteCount++
			respository.Db.Save(&userlike)
			respository.Db.Save(&video)
		}
		//获取redis中视频取消点赞信息保存进数据库
		for {
			if client.SCard("douyin:video:unliked:"+videoID).Val() == 0 {
				break
			}
			var userlike respository.UserLike
			userID := client.SPop(KeyVideoUnLikedSetPrefix + videoID).Val()

			userlike.VideoId, _ = strconv.ParseInt(videoID, 10, 64)
			userlike.LikeId, _ = strconv.ParseInt(userID, 10, 64)
			userlike.IsFavorite = false
			video.FavoriteCount--
			respository.Db.Save(&userlike)
			respository.Db.Save(&video)
		}

	}
}
