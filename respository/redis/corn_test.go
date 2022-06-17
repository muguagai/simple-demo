/******
** @创建时间 : 2022/6/17 22:20
** @作者 : MUGUAGAI
******/
package redis

import (
	"github.com/robfig/cron/v3"
	"testing"
)

func TestCorn(t *testing.T) {
	scheduler := cron.New()
	Cron(scheduler)
}
