package respository

import (
	"time"

	"gorm.io/gorm"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64 `json:"id,omitempty"`
	AuthorID      int64
	Author        User      `json:"author" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	CreateTime    time.Time `gorm:"column:create_time"`
	Title         string    `json:"title" gorm:"column:title"`
}

type Comment struct {
	Id         int64 `json:"id,omitempty"`
	UserID     int64
	User       User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VideoID    int64
	Video      Video  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:id"`
	Password      string `json:"password,omitempty" gorm:"column:password"`
	Name          string `json:"name,omitempty" gorm:"unique_index,column:name"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:followcount"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:followercount"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:IsFollow"`
	Token         string `gorm:"unique_index,column:token"`
}

type UserLike struct {
	gorm.Model
	LikeId     int64
	VideoId    int64
	IsFavorite bool
}

type FollowFollower struct {
	gorm.Model
	FollowId   int64
	FollowerId int64
	IsFavorite bool
}
