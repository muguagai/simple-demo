package middlewares

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/controller"

	"github.com/RaymondCode/simple-demo/util/jwt"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端URL携带Token
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		//authHeader := c.Request.Header.Get("Authorization")
		if token == "" {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "url缺少Auth Token")
			//中止函数
			c.Abort()
			return
		}

		_, err := jwt.ParseToken(token)
		if err != nil {
			fmt.Println(err)
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
