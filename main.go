package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"os"

	"github.com/RaymondCode/simple-demo/respository/redis"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/util"
)

func main() {
	c := cron.New()
	_, _ = c.AddFunc("@every 5s", redis.FavouriteToMysql)
	c.Start()
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
func Init() error {
	if err := respository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	if err := redis.InitClient(); err != nil {
		return err
	}
	config.GetAddress()
	//scheduler := cron.New()
	//开启定时任务
	return nil
}
