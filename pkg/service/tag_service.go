package service

import (
	"com.sj/admin/pkg/entity"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type ITagService interface {
	SaveUserTags(userId uint64, articleId uint64, labels string) bool
	SaveTag(tag *entity.ArticleTagPo) *entity.ArticleTag
}
type tagService struct {
	db                *gorm.DB
	articleTagService IArticleTagService
}

func NewTagService(articleTagService IArticleTagService, db *gorm.DB) ITagService {
	return &tagService{
		articleTagService: articleTagService,
		db:                db,
	}
}
func (t *tagService) SaveTag(tag *entity.ArticleTagPo) *entity.ArticleTag {
	var existData entity.ArticleTag
	result := t.db.Where("user_id = ? and name = ?", tag.UserId, tag.Name).First(&existData)
	if result.RowsAffected > 0 {
		updateResult := t.db.Select("UpdatedAt").Updates(&existData)
		if updateResult.RowsAffected > 0 {
			return &existData
		} else {
			logrus.Error(updateResult.Error)
			return nil
		}
	} else {
		copier.Copy(&existData, &tag)
		t.db.Create(&existData)
		return &existData
	}
}

func (t *tagService) SaveUserTags(userId uint64, articleId uint64, labels string) bool {
	if labels == "" {
		return false
	}
	labelSlice := strings.Split(labels, ",")
	if len(labelSlice) == 0 {
		return false
	}
	var tagIds []uint64
	for _, l := range labelSlice {
		articleTag := &entity.ArticleTagPo{
			UserId: userId,
			Name:   l,
		}
		tag := t.SaveTag(articleTag)
		if tag != nil {
			tagIds = append(tagIds, tag.ID)
		} else {
			return false
		}
	}
	ok := t.articleTagService.saveArticleTags(articleId, tagIds)
	return ok
}
