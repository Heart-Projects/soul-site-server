package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"errors"
	"github.com/jinzhu/copier"
)

type IArticleService interface {
	Save(po *entity.ArticlePo) (uint64, error)
	ListUserArticles(params *vo.UserArticlePageVo) (int64, []vo.UserArticleVo)

	GetArticleDetail(params *vo.ArticleDetailParams) (*vo.UserArticleVo, error)
}

type articleService struct {
	articleTagService IArticleTagService
	tagService        ITagService
	articleDao        dao.IArticleDao
}

func NewArticleService(articleTagService IArticleTagService, tagService ITagService, articleDao dao.IArticleDao) IArticleService {
	return &articleService{
		articleTagService: articleTagService,
		tagService:        tagService,
		articleDao:        articleDao,
	}

}
func (a *articleService) Save(po *entity.ArticlePo) (uint64, error) {
	var article entity.Article
	copier.Copy(&article, po)
	id := article.ID
	var ok bool
	var err error
	var dbArticle *entity.Article
	var optArticleId uint64
	dbError := dao.NewTransactionManager().Start(func(ctx *dao.TxContext) error {
		if id == 0 {
			ok, err, dbArticle = a.articleDao.Insert(&article, ctx)
		} else {
			ok, err, dbArticle = a.articleDao.UpdateById(&article, ctx)
		}
		// 处理文章标签
		if ok {
			tagOk := a.tagService.SaveUserTags(po.UserId, dbArticle.ID, po.Label)
			if tagOk {
				optArticleId = dbArticle.ID
				return nil
			} else {
				return errors.New("保存文章标签失败")
			}
		}
		return err
	})

	return optArticleId, dbError
}

func (a *articleService) ListUserArticles(params *vo.UserArticlePageVo) (int64, []vo.UserArticleVo) {
	total := a.articleDao.ListCount(params)
	if total == 0 {
		return total, []vo.UserArticleVo{}
	}
	dataList := a.articleDao.List(params)
	return total, dataList
}

func (a *articleService) GetArticleDetail(params *vo.ArticleDetailParams) (*vo.UserArticleVo, error) {
	r := a.articleDao.SelectById(params.ArticleId)
	if r == nil {
		return nil, errors.New("文章不存在")
	}
	r.Next = a.articleDao.SelectNextArticle(params.ArticleId, r.UserId)
	r.Pre = a.articleDao.SelectPreArticle(params.ArticleId, r.UserId)
	return r, nil
}
