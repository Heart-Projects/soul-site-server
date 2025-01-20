package controller

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type IArticleCategoryController interface {
	UserArticleCategory(c *gin.Context)
}
type ArticleCategoryController struct {
	articleCategorySrv service.IArticleCategoryService
}

func NewArticleCategoryController(articleCategorySrv service.IArticleCategoryService) IArticleCategoryController {
	return &ArticleCategoryController{
		articleCategorySrv: articleCategorySrv,
	}
}

func (a *ArticleCategoryController) UserArticleCategory(c *gin.Context) {
	_, u := utils.GetUserInfo(c)
	categoryList := a.articleCategorySrv.GetUserCategoryList(u.UserId)
	//transformArticleCategoryListToTreeData(categoryList)
	utils.DoResponseSuccessWithData(c, categoryList)
}

// 将分类的平铺结构转成树形结构
func transformArticleCategoryListToTreeData(list []entity.ArticleCategory) []*entity.TreeArticleCategory {
	size := len(list)
	if list == nil || size == 0 {
		return []*entity.TreeArticleCategory{}
	}

	treeData := make([]*entity.TreeArticleCategory, 0)
	dataMap := make(map[uint64]*entity.TreeArticleCategory, size)

	for _, value := range list {
		var treeItem entity.TreeArticleCategory
		copier.Copy(&treeItem, &value)
		treeItem.Children = make([]*entity.TreeArticleCategory, 0)
		dataMap[value.ID] = &treeItem
	}
	for _, value := range dataMap {
		if value.ParentId > 0 {
			parentNode := dataMap[value.ParentId]
			parentNode.Children = append(parentNode.Children, value)
		}
		if value.ParentId == 0 {
			treeData = append(treeData, value)
		}
	}
	return treeData
}
