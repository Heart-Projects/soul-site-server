package controller

import (
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IArticleColumnCategoryController interface {
	add(c *gin.Context)
	rename(c *gin.Context)
	remove(c *gin.Context)

	treeData(c *gin.Context)
}

type ArticleColumnCategoryController struct {
	server service.IArticleColumnCategoryService
}

func NewArticleColumnCategoryController(s service.IArticleColumnCategoryService) IArticleColumnCategoryController {
	return &ArticleColumnCategoryController{
		server: s,
	}
}

func (a *ArticleColumnCategoryController) add(c *gin.Context) {
	var params vo.ArticleColumnCategoryAddVo
	if err := c.ShouldBind(&params); err != nil {
		utils.DoResponseErrorMessage(c, "参数验证失败:"+err.Error())
		return
	}
	_, claims := utils.GetUserInfo(c)
	params.UserId = claims.UserId
	success, err, data := a.server.New(&params)

	if success {
		utils.DoResponseSuccessWithData(c, data)
	} else {
		utils.DoResponseErrorMessage(c, err.Error())
	}
}

func (a *ArticleColumnCategoryController) rename(c *gin.Context) {
	var params vo.ArticleColumnCategoryUpdateVo
	if err := c.ShouldBind(&params); err != nil {
		utils.DoResponseErrorMessage(c, "参数验证失败:"+err.Error())
		return
	}

	success, err := a.server.Rename(&params)
	if success {
		utils.DoResponseSuccess(c, "重命名成功")
	} else {
		utils.DoResponseErrorMessage(c, err.Error())
	}
}

func (a *ArticleColumnCategoryController) remove(c *gin.Context) {
	var idStr = c.Param("id")
	// 转换
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.DoResponseErrorMessage(c, "参数错误")
		return
	}
	success, err := a.server.Remove(id)
	if success {
		utils.DoResponseSuccess(c, "删除成功")
	} else {
		utils.DoResponseErrorMessage(c, err.Error())
	}

}

func (a *ArticleColumnCategoryController) treeData(c *gin.Context) {
	_, user := utils.GetUserInfo(c)
	categoryList := a.server.List(user.UserId)
	if len(categoryList) == 0 {
		utils.DoResponseSuccessWithData(c, []vo.ArticleColumnCategoryListVo{})
		return
	}
	voList := make([]*vo.ArticleColumnCategoryListVo, 0)
	// 注意v 在 range 循环中被重复使用
	for _, v := range categoryList {
		item := v
		voList = append(voList, &vo.ArticleColumnCategoryListVo{
			ArticleColumnCategory: &item,
			Children:              make([]utils.TreeItemFeature, 0),
		})
	}
	treeData := utils.TransformListToTreeData[*vo.ArticleColumnCategoryListVo](voList)
	utils.DoResponseSuccessWithData(c, treeData)
}
