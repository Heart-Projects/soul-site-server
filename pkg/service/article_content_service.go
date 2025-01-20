package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
)

type IArticleContentService interface {
	GetOne(source entity.ArticleSource, docId uint64) (bool, *entity.ArticleContent)

	SaveContent(source entity.ArticleSource, docId uint64, content string) (bool, error)
}

type articleContentService struct {
	contentDao dao.IArticleContentDao
}

func NewArticleContentService(contentDao dao.IArticleContentDao) IArticleContentService {
	return &articleContentService{
		contentDao: contentDao,
	}
}

func (a *articleContentService) GetOne(source entity.ArticleSource, docId uint64) (bool, *entity.ArticleContent) {
	return a.contentDao.SelectSourceItem(source, docId)
}

func (a *articleContentService) SaveContent(source entity.ArticleSource, docId uint64, content string) (bool, error) {
	return a.contentDao.SaveSourceItem(source, docId, content, nil)
}
