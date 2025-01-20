package controller

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ISiteCategoryController interface {
	listAllSiteCategory(c *gin.Context)
}

type siteCategoryController struct {
	siteCategoryDao dao.ISiteCategoryDao
}

func NewSiteCategoryController(categoryDao dao.ISiteCategoryDao) ISiteCategoryController {
	return &siteCategoryController{siteCategoryDao: categoryDao}
}

// 获取站点分类
// @Summary      获取站点分类
// @Description  获取全部有效的分类
// @Tags         site
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.TreeSiteCategory
// @Router       /site-category/list [get]
func (s *siteCategoryController) listAllSiteCategory(c *gin.Context) {
	categories := s.siteCategoryDao.SelectAllList()
	//treeData := transformSiteCategoryListToTreeData(categories)
	utils.DoResponseSuccessWithData(c, categories)
}

// 将分类的平铺结构转成树形结构
func transformSiteCategoryListToTreeData(list []entity.SiteCategory) []*entity.TreeSiteCategory {
	size := len(list)
	if list == nil || size == 0 {
		return []*entity.TreeSiteCategory{}
	}

	treeData := make([]*entity.TreeSiteCategory, 0)
	dataMap := make(map[uint64]*entity.TreeSiteCategory, size)

	for _, value := range list {
		var treeItem entity.TreeSiteCategory
		copier.Copy(&treeItem, &value)
		treeItem.Children = make([]*entity.TreeSiteCategory, 0)
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
