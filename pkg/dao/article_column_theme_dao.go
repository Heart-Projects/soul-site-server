package dao

import (
	"com.sj/admin/pkg/entity"
	"gorm.io/gorm"
)

// IArticleColumnThemeDao 专栏主题
type IArticleColumnThemeDao interface {
	IBaseDao[entity.ArticleColumnTheme]
}

type articleColumnThemeDao struct {
	db *gorm.DB
	BaseDao[entity.ArticleColumnTheme]
}

func newArticleColumnThemeDao(db *gorm.DB) IArticleColumnThemeDao {
	return &articleColumnThemeDao{db: db}
}

func (a *articleColumnThemeDao) Insert(E *entity.ArticleColumnTheme, txContext *TxContext) (bool, error, uint64) {
	return true, nil, 1
}

func (a *articleColumnThemeDao) Update(E *entity.ArticleColumnTheme, txContext *TxContext) (bool, error) {
	return true, nil
}

func (a *articleColumnThemeDao) Delete(E *entity.ArticleColumnTheme, txContext *TxContext) (bool, error) {
	return true, nil
}

func (a *articleColumnThemeDao) SelectOne(id uint64) *entity.ArticleColumnTheme {
	var c entity.ArticleColumnTheme
	a.db.Where("id = ?", id).Find(&c)
	return &c
}

func (a *articleColumnThemeDao) SelectList(conditions ...interface{}) []entity.ArticleColumnTheme {
	var c []entity.ArticleColumnTheme
	a.db.Where(conditions).Find(&c)
	return c
}
