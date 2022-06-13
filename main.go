package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"os"

	"github.com/RaymondCode/simple-demo/respository/redis"

	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

func main() {
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
	return nil
}
