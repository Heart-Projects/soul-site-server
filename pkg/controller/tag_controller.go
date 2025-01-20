package controller

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ITagController interface {
	addTag(c *gin.Context)
}
type TagController struct {
	tagSrv service.ITagService
}

func NewTagController(tagSrv service.ITagService) ITagController {
	return &TagController{
		tagSrv: tagSrv,
	}
}

func (t *TagController) addTag(c *gin.Context) {
	var articleTagPo entity.ArticleTagPo
	_, user := utils.GetUserInfo(c)
	err := c.ShouldBind(&articleTagPo)
	if err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	articleTagPo.UserId = user.UserId
	tag := t.tagSrv.SaveTag(&articleTagPo)
	utils.DoResponseWithConditionMessage(c, tag != nil, "k", tag)
}
