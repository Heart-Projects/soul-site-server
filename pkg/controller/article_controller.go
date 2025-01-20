package controller

import (
	"com.sj/admin/pkg/constant"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type IArticleController interface {
	saveArticle(c *gin.Context)

	listUserArticles(c *gin.Context)

	articleDetail(c *gin.Context)

	articleItemDetail(c *gin.Context)
}
type ArticleController struct {
	articleSrv service.IArticleService
}

func NewArticleController(articleSrv service.IArticleService) IArticleController {
	return &ArticleController{
		articleSrv: articleSrv,
	}
}

func (a *ArticleController) saveArticle(c *gin.Context) {
	var article entity.ArticlePo
	if err := c.ShouldBind(&article); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	_, u := utils.GetUserInfo(c)
	article.UserId = u.UserId
	result, err := a.articleSrv.Save(&article)
	logrus.Info("save article is ", result)
	utils.DoResponseWithErrorMessage(c, err, "保存文章失败", result)
}

func (a *ArticleController) listUserArticles(c *gin.Context) {
	var articlePageVo vo.UserArticlePageVo
	if err := c.ShouldBind(&articlePageVo); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}

	if articlePageVo.PageSize <= 0 {
		articlePageVo.PageSize = constant.PAGE_SIZE
	}

	if articlePageVo.PageIndex <= 0 {
		articlePageVo.PageIndex = 1
	}

	_, u := utils.GetUserInfo(c)
	articlePageVo.UserId = u.UserId
	total, list := a.articleSrv.ListUserArticles(&articlePageVo)
	articlePageVo.List = list
	articlePageVo.Total = total
	utils.DoResponseSuccessWithData(c, articlePageVo)
}

func (a *ArticleController) articleDetail(c *gin.Context) {
	var params vo.ArticleDetailParams
	if err := c.ShouldBindUri(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	result, err := a.articleSrv.GetArticleDetail(&params)
	if err != nil {
		utils.DoResponseWithErrorMessage(c, err, "获取文章详情失败", result)
	}
	utils.DoResponseSuccessWithData(c, result)
}

func (a *ArticleController) articleItemDetail(c *gin.Context) {
	articleIdStr := c.Param("articleId")
	articleId, err := strconv.ParseUint(articleIdStr, 10, 64)
	if articleIdStr == "" || err != nil {
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}

	result, err := a.articleSrv.GetArticleDetail(&vo.ArticleDetailParams{ArticleId: articleId})
	if err != nil {
		utils.DoResponseWithErrorMessage(c, err, "获取文章详情失败", result)
	}
	utils.DoResponseSuccessWithData(c, result)
}
