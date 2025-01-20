package controller

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IArticleContentController interface {

	// Save 保存文章内容
	Save(c *gin.Context)

	// Content 获取文章内容
	Content(c *gin.Context)
}

type ArticleContentController struct {
	srv service.IArticleContentService
}

func NewArticleContentController(srv service.IArticleContentService) IArticleContentController {
	return &ArticleContentController{
		srv: srv,
	}
}

func (a *ArticleContentController) Save(c *gin.Context) {
	var params vo.ArticleContentSaveVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, err := a.srv.SaveContent(entity.ArticleSource(params.Source), params.DocId, params.Content)
	utils.DoResponseWithCondition(c, success, nil, err)
}

func (a *ArticleContentController) Content(c *gin.Context) {
	var params vo.ArticleContentRequestVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}

	if params.DocId == 0 {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败, docId: %d", params.DocId))
		return
	}

	success, data := a.srv.GetOne(entity.ArticleSource(params.Source), params.DocId)

	utils.DoResponseWithCondition(c, success, data, errors.New("文章不存在"))
}
