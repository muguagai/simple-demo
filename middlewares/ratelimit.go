package middlewares

import (
	"github.com/RaymondCode/simple-demo/respository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 创建指定填充速率和容量大小的令牌桶
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusOK, respository.Response{StatusCode: 1, StatusMsg: "登录次数过于频繁，请稍后重试"})
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
