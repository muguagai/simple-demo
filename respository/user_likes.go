package respository

import "sync"

type UserLikeDao struct {
}

var userLikeDao *UserLikeDao
var userLikeOnce sync.Once

func NewUserLikeDaoInstance() *UserLikeDao {
	userOnce.Do(
		func() {
			userLikeDao = &UserLikeDao{}
		})
	return userLikeDao
}

//根据userid查询用户喜爱列表
func (*UserLikeDao) QueryFavoriteListByUserId(userId int64) []Video {
	var Videos []Video = nil
	var videoIds []*int64
	Db.Table("user_likes").Select("video_id").Where("like_id = ?", userId).Where("is_favorite= ?", 1).Find(&videoIds)
	if len(videoIds) == 0 {
		return nil
	}
	Db.Find(&Videos, videoIds)
	return Videos
}

func (*UserLikeDao) QueryUserLikeByVideoIDandLikeId(videoId int64, likeId int64) (UserLike, bool) {
	var userLike UserLike
	find := Db.Table("user_likes").Where("video_id = ?", videoId).Where("like_id = ?", likeId).Find(&userLike)
	if find != nil {
		return userLike, true
	}

	return userLike, false
}
