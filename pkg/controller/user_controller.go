package controller

import (
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IUserController interface {
	login(c *gin.Context)
	getUserByToken(c *gin.Context)
	logout(ctx *gin.Context)

	userArticleSummaryInfo(ctx *gin.Context)
}

type userController struct {
	userSrv service.IUserService
}

func NewUserController(userSrv service.IUserService) IUserController {
	return &userController{
		userSrv: userSrv,
	}
}

// 登陆
func (u *userController) login(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	loginSuccess, err, user := u.userSrv.CheckLogin(username, password)
	logrus.Info("login result is ", loginSuccess)
	if loginSuccess {
		token, err := utils.GenerateJWT(user)

		utils.DoResponseWithError(c, err, token)
	} else {
		utils.DoResponseError(c, err)
	}
}

func (u *userController) logout(c *gin.Context) {
	utils.DoResponseSuccess(c, "退出成功")
}

// 获取用户信息
func (u *userController) getUserByToken(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		logrus.Error("can not get access token from header, the result is ", token)
		utils.DoResponseErrorMessage(c, "无法获取token 信息")
	}
	claims, err := utils.ParseToken(token)
	utils.DoResponseWithConditionMessage(c, err == nil, "token 解析成功", claims)
}

func (u *userController) userArticleSummaryInfo(ctx *gin.Context) {
	_, user := utils.GetUserInfo(ctx)
	ur := u.userSrv.ArticleSummaryData(user.UserId)
	if ur == nil {
		utils.DoResponseErrorMessage(ctx, "获取用户文章统计数据失败")
		return
	}
	utils.DoResponseSuccessWithData(ctx, ur)
}
