package jwt

import (
	"AI_Server/init/conf"
	"AI_Server/internal/modeles"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// PayLoad jwt中payload数据
type PayLoad struct {
	UserId uint             `json:"userId"`
	Role   modeles.UserRole `json:"role"` // 权限  1 管理员 2 普通用户
}

type CustomClaims struct {
	PayLoad
	jwt.RegisteredClaims
}

// GenToken 创建 Token
func GenToken(payload PayLoad) (string, error) {
	j := conf.GlobalConfig.Jwt
	claim := CustomClaims{
		PayLoad: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Expires)),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(j.Secret))
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GlobalConfig.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
