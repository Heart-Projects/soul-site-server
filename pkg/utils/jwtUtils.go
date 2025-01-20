package utils

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/types"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"time"
)

var TokenSecret []byte = []byte("sdkdkksdkskdkskdk")

// ParseToken  反解析jwt token
func ParseToken(token string) (*types.UserJwtClaims, error) {
	claims, err := jwt.ParseWithClaims(token, &types.UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, jwt.NewValidationError("token 解析失败", 401)
	}
	if claims.Valid {
		userJwtClaims := claims.Claims.(*types.UserJwtClaims)
		return userJwtClaims, nil
	} else {
		return nil, jwt.NewValidationError("token 解析失败", 401)
	}
}

// GenerateJWT  生成jwt token
func GenerateJWT(user *entity.SysUser) (string, error) {
	var u = types.UserJwtClaims{
		Username: user.Name,
		UserId:   user.ID,
		Roles:    []string{"admin", "admin2"},
		Menus:    []string{"aa"},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	signedString, err := claims.SignedString(TokenSecret)
	if err == nil {
		return signedString, nil
	} else {
		logrus.Error("签名失败, 原因为:" + err.Error())
		return "", errors.New("签名失败, 原因为:" + err.Error())
	}
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) (bool, *types.UserJwtClaims) {
	value, exists := c.Get("_$user")
	if exists {
		user := value.(*types.UserJwtClaims)
		return true, user
	} else {
		return false, nil
	}
}
