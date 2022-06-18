package jwt

import (
	"errors"
	"github.com/RaymondCode/simple-demo/respository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Password string `json:"password"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//定义Secret
var mySecret = []byte("夏天夏天悄悄过去")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

//定义JWT的过期时间
const TokenExpireDuration = time.Hour * 1000000

// GenToken 生成access token 和 refresh token
func GenToken(username string, password string) (aToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		password, // 自定义字段
		username, // 自定义字段
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "muguagai",                                 // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return
}

//TODO 解析JWT

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	respository.UsersLoginInfo[claims.Username] = respository.NewUserDaoInstance().QueryUserByUserName(claims.Username)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}
