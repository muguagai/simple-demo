package respository

//根据followid查询用户关注列表
func QueryFollowListByUserId(userId int64) []User {
	var Users []User = nil
	var userIds []*int64
	Db.Table("follow_followers").Select("follow_id").Where("follower_id = ?", userId).Where("is_favorite= ?", 1).Find(&userIds)
	if len(userIds) == 0 {
		return nil
	}
	Db.Find(&Users, userIds)
	return Users
}

//根据followid查询用户粉丝列表
func QueryFollowerListByUserId(userId int64) []User {
	var Users []User = nil
	var userIds []*int64
	Db.Table("follow_followers").Select("follower_id").Where("follow_id = ?", userId).Where("is_favorite= ?", 1).Find(&userIds)
	if len(userIds) == 0 {
		return nil
	}
	Db.Find(&Users, userIds)
	return Users
}
