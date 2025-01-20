package dao

import (
	"com.sj/admin/pkg/entity"
	"gorm.io/gorm"
)

type IArticleCategoryDao interface {
	SelectList(userId uint64) []entity.ArticleCategory
}
type articleCategoryDao struct {
	db *gorm.DB
}

func newArticleCategoryDao(db *gorm.DB) IArticleCategoryDao {
	return &articleCategoryDao{db: db}
}

// SelectList 查询用户所有的标签
func (a *articleCategoryDao) SelectList(userId uint64) []entity.ArticleCategory {
	var c []entity.ArticleCategory
	a.db.Where("user_id = ?", userId).Find(&c)
	return c
}
