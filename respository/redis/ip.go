/******
** @创建时间 : 2022/6/18 23:36
** @作者 : MUGUAGAI
******/
package redis

import (
	"strconv"
	"time"
)

func QueryIpCount(ip string) int64 {
	KEY := KeyIpCount + ip
	//根据KEY 查找ip是否存在
	val := client.Exists(KEY).Val()
	//存在则将value+1
	if val != 0 {
		client.Incr(KEY)
		get := client.Get(KEY).Val()
		count, _ := strconv.ParseInt(get, 10, 64)
		return count
	} else {
		//不存在则创建KEY
		client.Set(KEY, 1, time.Second*60)
		return 1
	}
}
