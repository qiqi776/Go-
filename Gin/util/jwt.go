package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 用于签名的字符串
var mySigningKey = []byte("qifeng.com")

type MyCustomClaims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 使用你的自定义声明创建 jwt
func GenToken(username string) (string, error) {
	// 创建Claims
	claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "qifeng", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 解析token
	// 注意：我们这里用 ParseWithClaims 和我们自定义的 MyCustomClaims 结构体
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token 并且转换 claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
