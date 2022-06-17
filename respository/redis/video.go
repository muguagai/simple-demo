/******
** @创建时间 : 2022/6/1 11:21
** @作者 : MUGUAGAI
******/
package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	LikedScore       float64 = 432 // 每一个点赞加432分
)

func LikedForVideo(videoID string, isFavorite bool, userID int64) (err error) {
	//uid := strconv.FormatInt(userID, 10)
	pipeline := client.TxPipeline()
	var op float64
	key := KeyUserLikedHashPrefix
	videokey := KeyVideoLikedSetPrefix + videoID
	if isFavorite == true {
		//将videoid添加到用户喜爱列表
		op = 1
		pipeline.SAdd(key, videoID)
		pipeline.SAdd(videokey, userID)
	} else {
		op = -1
		pipeline.SRem(videokey, userID)
	}

	ret, err := pipeline.ZRevRangeWithScores(KeyVideoScoreZSet, 0, 0).Result()
	fmt.Println(ret)
	score, err := pipeline.ZIncrBy(KeyVideoScoreZSet, LikedScore*op, videoID).Result() // 更新分数
	fmt.Println(score)
	if err != nil {
		return err
	}
	_, err = pipeline.Exec()
	return err
}

//按照分数大小顺序查询视频的IDS
func GetIDsFormKey(start int64) (IDs []string, err error, end int64) {
	result, err := client.ZRevRange(KeyVideoScoreZSet, start, start+29).Result()
	//视频已经全部浏览完
	if len(result) == 0 {
		start = 0
		result, err = client.ZRevRange(KeyVideoScoreZSet, 0, 29).Result()
	}
	if len(result) < 30 {
		end = start + int64(len(result))
		return result, err, end
	}
	return result, err, start + 30
}

// CreatePost 使用hash存储视频信息
func CreateVideo(videoID int64) (err error) {
	now := float64(time.Now().Unix())
	//likedKey := KeyVideoLikedZSetPrefix + strconv.Itoa(int(videoID))
	//	videoInfo := map[string]interface{}{
	//		"video:id": videoID,
	//		"time":     now,
	//	}
	pipeline := client.TxPipeline()
	//pipeline.Expire(likedKey, time.Second*OneWeekInSeconds*10000) // 100年不过期
	//pipeline.HMSet(likedKey, videoInfo)
	num, err := pipeline.ZAdd(KeyVideoScoreZSet, redis.Z{ // 添加到分数的ZSet
		Score:  now,
		Member: videoID,
	}).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)
	pipeline.ZAdd(KeyVideoTimeZSet, redis.Z{ // 添加到时间的ZSet
		Score:  now,
		Member: videoID,
	})
	_, err = pipeline.Exec()
	return
}

func GetFavouriteVideo(userID int64) (videoIDS []string) {
	videoIDS = client.SMembers(KeyUserLikedHashPrefix + strconv.Itoa(int(userID))).Val()
	return videoIDS
}
