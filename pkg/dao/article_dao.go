package dao

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IArticleDao interface {
	Insert(article *entity.Article, txContext *TxContext) (bool, error, *entity.Article)
	UpdateById(article *entity.Article, txContext *TxContext) (bool, error, *entity.Article)

	ListCount(param *vo.UserArticlePageVo) int64

	List(param *vo.UserArticlePageVo) []vo.UserArticleVo

	SelectById(id uint64) *vo.UserArticleVo

	SelectNextArticle(currentArticleId uint64, userId uint64) *vo.UserNavArticleVo

	SelectPreArticle(currentArticleId uint64, userId uint64) *vo.UserNavArticleVo
}

type articleDao struct {
	db *gorm.DB
}

func newArticleDao(db *gorm.DB) IArticleDao {
	return &articleDao{db: db}
}

func (a *articleDao) Insert(article *entity.Article, txContext *TxContext) (bool, error, *entity.Article) {
	result := FindTxDb(a.db, txContext).Create(article)
	if success := result.RowsAffected > 0; success {
		return success, nil, article
	}
	return false, result.Error, nil
}

func (a *articleDao) UpdateById(article *entity.Article, txContext *TxContext) (bool, error, *entity.Article) {
	if article.ID <= 0 {
		logrus.Error("文章ID 为空")
		return false, errors.New("文章ID 为空"), nil
	}
	result := FindTxDb(a.db, txContext).Model(article).Select("*").Omit("CreatedAt", "UserId").Updates(article)
	return result.RowsAffected > 0, result.Error, article
}

func (a *articleDao) ListCount(param *vo.UserArticlePageVo) int64 {
	var total int64
	a.db.Model(&entity.Article{}).Where("user_id = ?", param.UserId).Count(&total)
	return total
}

func (a *articleDao) List(param *vo.UserArticlePageVo) []vo.UserArticleVo {
	var list []entity.Article
	a.db.Where("user_id = ?", param.UserId).Limit(param.PageSize).Offset((param.PageIndex - 1) * param.PageSize).Order("created_at desc").Find(&list)
	var articleList = make([]vo.UserArticleVo, 0)
	if len(list) > 0 {
		for _, v := range list {
			var labels []entity.ArticleTag
			a.db.Model(&entity.ArticleTag{}).Joins(" inner join article_tag_relation at on at.tag_id = article_tag.id ").Where("at.article_id = ?", v.ID).Find(&labels)
			articleList = append(articleList, vo.UserArticleVo{Labels: &labels, Article: v})
		}

	}
	return articleList
}

func (a *articleDao) SelectById(id uint64) *vo.UserArticleVo {
	var article entity.Article
	result := a.db.Where("id = ?", id).Take(&article)
	if result.RowsAffected == 0 {
		return nil
	}
	var labels []entity.ArticleTag
	a.db.Model(&entity.ArticleTag{}).Joins(" inner join article_tag_relation at on at.tag_id = article_tag.id ").Where("at.article_id = ?", id).Find(&labels)
	return &vo.UserArticleVo{Labels: &labels, Article: article}
}

func (a *articleDao) SelectNextArticle(currentArticleId uint64, userId uint64) *vo.UserNavArticleVo {
	var nextArticle entity.Article
	a.db.Where("id > ? and user_id = ? and status = ?", currentArticleId, userId, entity.ArticleStatusPublish).Order("id asc").Take(&nextArticle)
	if nextArticle.ID == 0 {
		return nil
	}
	return &vo.UserNavArticleVo{ID: nextArticle.ID, Title: nextArticle.Title}
}

func (a *articleDao) SelectPreArticle(currentArticleId uint64, userId uint64) *vo.UserNavArticleVo {
	var preArticle entity.Article
	a.db.Where("id < ? and user_id = ? and status = ?", currentArticleId, userId, entity.ArticleStatusPublish).Order("id desc").Take(&preArticle)
	if preArticle.ID == 0 {
		return nil
	}
	return &vo.UserNavArticleVo{ID: preArticle.ID, Title: preArticle.Title}
}
