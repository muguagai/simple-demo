/******
** @创建时间 : 2022/6/17 19:01
** @作者 : MUGUAGAI
******/
package service

import (
	"github.com/RaymondCode/simple-demo/respository"
	"gorm.io/gorm"
)

func RelationAction(follow *respository.User, actiontype string, follower respository.User, follow_follower respository.FollowFollower, find *gorm.DB) (err error) {
	if actiontype == "1" {
		follow.FollowCount++
		follower.FollowerCount++

		respository.NewUserDaoInstance().SaveUser(follower)
		respository.NewUserDaoInstance().SaveUser(*follow)
		if find != nil {
			follow_follower.FollowId = follow.Id
			follow_follower.FollowerId = follower.Id
			follow_follower.IsFavorite = true
			respository.Db.Save(&follow_follower)
		} else {
			respository.Db.Create(&follow_follower)
		}
	}
	if actiontype == "2" {
		follow.FollowCount--
		follower.FollowerCount--
		respository.NewUserDaoInstance().SaveUser(follower)
		respository.NewUserDaoInstance().SaveUser(*follow)
		follow_follower.FollowId = follow.Id
		follow_follower.FollowerId = follower.Id
		follow_follower.IsFavorite = false
		respository.Db.Save(&follow_follower)
	}
	return nil
}
