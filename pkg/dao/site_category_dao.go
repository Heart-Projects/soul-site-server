package dao

import (
	"com.sj/admin/pkg/entity"
	"gorm.io/gorm"
)

type ISiteCategoryDao interface {
	SelectAllList() []entity.SiteCategory
}

type siteCategoryDao struct {
	db *gorm.DB
}

func newSiteCategoryDao(db *gorm.DB) ISiteCategoryDao {
	return &siteCategoryDao{db: db}
}

func (s *siteCategoryDao) SelectAllList() []entity.SiteCategory {
	var siteCategories []entity.SiteCategory
	s.db.Find(&siteCategories)
	return siteCategories
}
