/******
** @创建时间 : 2022/6/1 11:19
** @作者 : MUGUAGAI
******/
package redis

const (
	KeyVideoInfoHashPrefix = "douyin:video:"
	KeyVideoTimeZSet       = "douyin:video:time"   // zset;帖子及发帖时间定义
	KeyVideoScoreZSet      = "douyin:video:score"  // zset;帖子及投票分数定义
	KeyVideoLikedSetPrefix = "douyin:video:liked:" // set;记录视频点赞用户id;
	KeyUserLikedHashPrefix = "douyin:user:liked:"  //set;记录用户id的视频点赞;

	//KeyCommunityPostSetPrefix = "bluebell:community:" // set保存每个分区下帖子的id
)
