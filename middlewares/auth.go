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
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		//authHeader := c.Request.Header.Get("Authorization")
		if token == "" {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "请求头缺少Auth Token")
			//中止函数
			c.Abort()
			return
		}
		// 按空格分割
		/*parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}*/
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		_, err := jwt.ParseToken(token)
		if err != nil {
			fmt.Println(err)
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		//c.Set(controller.ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
