/******
** @创建时间 : 2022/6/3 22:31
** @作者 : MUGUAGAI
******/
package service

import (
	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/respository/redis"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/RaymondCode/simple-demo/util/jwt"
	"go.uber.org/zap"
)

func PublishVideo(video respository.Video) (err error) {
	respository.Db.Create(&video)
	err = redis.CreateVideo(video.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetVideoList(start int64, token string) (videos []respository.Video, end int64) {
	//从REDIS获取视频ID信息
	ids, err, End := redis.GetIDsFormKey(start)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	IDS := util.String2Int(ids)
	//根据ID查询视频列表并且按照IDS的排序查询
	Videos := respository.GetVideoListByIDs(IDS)
	for i := 0; i < len(Videos); i++ {
		var user respository.User
		respository.Db.Where("id = ?", Videos[i].AuthorID).Find(&user)
		Videos[i].Author = user
	}
	parseToken, _ := jwt.ParseToken(token)
	username := parseToken.Username
	LoginUser := respository.UsersLoginInfo[username]
	//登录状态下将作者与用户关注状态和视频点赞状态关联
	if len(token) != 0 {
		for i := 0; i < len(Videos); i++ {
			var follow_follower respository.FollowFollower
			find := respository.Db.Table("follow_followers").Where("follow_id = ?", Videos[i].AuthorID).Where("follower_id = ?", LoginUser.Id).Find(&follow_follower)
			if find != nil {
				Videos[i].Author.IsFollow = follow_follower.IsFavorite
			}
		}
		for j := 0; j < len(Videos); j++ {
			var userLike respository.UserLike
			find := respository.Db.Table("user_likes").Where("like_id = ?", LoginUser.Id).Where("video_id=?", Videos[j].Id).Find(&userLike)
			if find != nil {
				Videos[j].IsFavorite = userLike.IsFavorite
			}
		}
	}

	return Videos, End
}
