package respository

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	Init()
	/*Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Video{})
	Db.AutoMigrate(&UserLike{})
	Db.AutoMigrate(&Comment{})
	Db.AutoMigrate(&FollowFollower{})*/
	//user, _ := userDao.QueryUserById(537331302576164864)
	//fmt.Println(user)
	var userlike UserLike
	userlike.VideoId = 1
	userlike.LikeId = 1
	Db.Save(&userlike)

	iDs := GetVideoListByIDs([]int64{539237567535517696, 539239326517563392})
	fmt.Println(iDs)
}
