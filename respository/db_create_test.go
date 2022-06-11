package respository

import (
	"testing"
)

func TestCreateDB(t *testing.T) {
	Init()
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Video{})
	Db.AutoMigrate(&UserLike{})
	Db.AutoMigrate(&Comment{})
	Db.AutoMigrate(&FollowFollower{})
}
