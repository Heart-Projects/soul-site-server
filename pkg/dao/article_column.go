package dao

import (
	"com.sj/admin/pkg/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// IArticleColumnDao 专栏
type IArticleColumnDao interface {
	IBaseDao[entity.ArticleColumn]

	SelectOneByField(field string, value interface{}) *entity.ArticleColumn
}

type articleColumnDao struct {
	db *gorm.DB
	BaseDao[entity.ArticleColumn]
}

func newArticleColumnDao(db *gorm.DB) IArticleColumnDao {
	return &articleColumnDao{db: db}
}

func (a *articleColumnDao) Insert(e *entity.ArticleColumn, txContext *TxContext) (bool, error, uint64) {
	r := a.FindDb(txContext, a.db).Create(e)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil, e.ID
	}
	return false, r.Error, 0
}

func (a *articleColumnDao) Update(e *entity.ArticleColumn, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Model(e).Updates(e)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil
	}
	return false, r.Error
}

func (a *articleColumnDao) Delete(e *entity.ArticleColumn, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Delete(e)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil
	}
	return false, errors.New(fmt.Sprintf("Column not exist: %d", e.ID))
}

func (a *articleColumnDao) SelectOne(id uint64) *entity.ArticleColumn {
	var c entity.ArticleColumn
	a.db.Where("id = ?", id).Find(&c)
	return &c
}

func (a *articleColumnDao) SelectOneByField(field string, value interface{}) *entity.ArticleColumn {
	var c entity.ArticleColumn
	a.db.Where(fmt.Sprintf("%s = ?", field), value).Find(&c)
	return &c
}

func (a *articleColumnDao) SelectList(conditions ...interface{}) []entity.ArticleColumn {
	var c []entity.ArticleColumn
	a.db.Where("category_id = ? ", conditions[0]).Order("created_at desc").Find(&c)
	return c
}
