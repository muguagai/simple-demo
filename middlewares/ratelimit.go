package middlewares

import (
	"github.com/RaymondCode/simple-demo/respository"
	"github.com/RaymondCode/simple-demo/respository/redis"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var IpCallCount map[int64]string

// 创建指定填充速率和容量大小的令牌桶
func RateLimitLoginMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit..
		ip := c.ClientIP()
		//判断IP访问次数
		count := redis.QueryIpCount(ip)
		if count > 100 {
			c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "登录次数过于频繁，请稍后重试"})
			c.Abort()
			return
		}
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "登录次数过于频繁，请稍后重试"})
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
func RateLimitFavouriteMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "点赞次数过于频繁，请稍后重试"})
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
