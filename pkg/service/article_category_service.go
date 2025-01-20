package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
)

type IArticleCategoryService interface {
	GetUserCategoryList(userId uint64) []entity.ArticleCategory
}

type articleCategoryService struct {
	articleCategoryDao dao.IArticleCategoryDao
}

// 构造一个对象
func newArticleCategory(articleCategoryDao dao.IArticleCategoryDao) IArticleCategoryService {
	return &articleCategoryService{
		articleCategoryDao: articleCategoryDao,
	}
}

func (a *articleCategoryService) GetUserCategoryList(userId uint64) []entity.ArticleCategory {
	return a.articleCategoryDao.SelectList(userId)
}
