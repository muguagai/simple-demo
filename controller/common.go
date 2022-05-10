package controller

import (
	"gorm.io/gorm"
	"time"
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
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:id"`
	Password      string `json:"password,omitempty" gorm:"column:password"`
	Name          string `json:"name,omitempty" gorm:"column:name"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:followcount"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:followercount"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:IsFollow"`
}

type UserLike struct {
	gorm.Model
	LikeId     int64
	VideoId    int64
	IsFavorite bool
}
