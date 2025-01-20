package middleware

import (
	"com.sj/admin/pkg/model"
	"com.sj/admin/pkg/utils"
	"com.sj/admin/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var safeUrl = []string{"/user/login", "/user/logout", "/swagger/*any", "/api-docs/*any", "/favicon.ico", "/", "/debug/pprof/*", "/file/upload"}

func Auth(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.RequestURI
		logrus.Info("path is" + path)
		if checkIsSafeUrl(path) {
			logrus.Info("path " + path + "is safe url, will pass")
			c.Next()
		} else if hasLogin, userJwtClaims := checkHasEffectiveLogin(c); hasLogin {
			logrus.Info("path " + path + " is not safe url, check auth")
			// 将用户信息放入上下文中
			c.Set("_$user", userJwtClaims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, model.Response.FailWithMsg("无权限操作，请先登陆"))
		}
	}
}

func checkIsSafeUrl(path string) bool {
	result := false
	for i := range safeUrl {
		if safeUrl[i] == path {
			result = true
			break
		}
	}
	return result
}

func checkHasEffectiveLogin(c *gin.Context) (bool, *types.UserJwtClaims) {
	requestURI := c.Request.RequestURI
	token := c.GetHeader("Authorization")
	logrus.WithFields(logrus.Fields{"requestURI": requestURI, "token": token}).Info("request uri is token is ")
	s, _ := c.Cookie("access_token")
	logrus.WithFields(logrus.Fields{"requestURI": requestURI, "s": s}).Info("request uri is token is ")
	if token == "" {
		return false, nil
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		logrus.Error(err.Error())
		return false, nil
	} else {
		return true, claims
	}
}
