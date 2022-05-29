package respository

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	Init()
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Video{})
	Db.AutoMigrate(&UserLike{})
	Db.AutoMigrate(&Comment{})
	Db.AutoMigrate(&FollowFollower{})
	user, _ := userDao.QueryUserById(537331302576164864)
	fmt.Println(user)
	/*var video = Video{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.5:8080/static/bear.mp4",
		CoverUrl:      "http://192.168.1.5:8080/static/屏幕截图 2021-02-16 163146.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}*/
}
