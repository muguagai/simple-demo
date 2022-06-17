/******
** @创建时间 : 2022/6/17 22:19
** @作者 : MUGUAGAI
******/
package redis

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func Cron(c *cron.Cron) {
	i := 1
	EntryID, err := c.AddFunc("*/1 * * * *", func() {
		FavouriteToMysql()
		fmt.Println(time.Now(), "每分钟执行一次", i)
		i++
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
	time.Sleep(time.Minute * 5)
}
