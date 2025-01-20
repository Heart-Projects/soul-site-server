package service

import (
	"com.sj/admin/pkg/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IArticleTagService interface {
	saveArticleTags(articleId uint64, tagIds []uint64) bool
}

type articleTagService struct {
	db *gorm.DB
}

func NewArticleTagService(db *gorm.DB) IArticleTagService {
	return &articleTagService{
		db: db,
	}
}
func (a *articleTagService) saveArticleTags(articleId uint64, tagIds []uint64) bool {
	// 删除已经存在的数据
	tx := a.db.Where("article_id=?", articleId).Delete(&entity.ArticleTagRelation{})
	logrus.WithField("rows", tx.RowsAffected).Debug()
	// 查询新的数据
	for _, tagId := range tagIds {
		at := &entity.ArticleTagRelation{ArticleId: articleId, TagId: tagId}
		r := a.db.Create(at)
		if r.Error != nil {
			return false
		}
	}
	return true
}
