package dao

import (
	"com.sj/admin/pkg/entity"
	"gorm.io/gorm"
)

type IArticleContentDao interface {
	IBaseDao[entity.ArticleContent]

	SaveSourceItem(source entity.ArticleSource, docId uint64, content string, txContext *TxContext) (bool, error)

	DeleteSourceItem(source entity.ArticleSource, docId uint64, txContext *TxContext) (bool, error)

	SelectSourceItem(source entity.ArticleSource, docId uint64) (bool, *entity.ArticleContent)
}

type articleContentDao struct {
	db *gorm.DB
	BaseDao[entity.ArticleContent]
}

func newArticleContentDao(db *gorm.DB) IArticleContentDao {
	return &articleContentDao{db: db}
}

func (a *articleContentDao) Insert(e *entity.ArticleContent, txContext *TxContext) (bool, error, uint64) {
	r := a.FindDb(txContext, a.db).Create(e)
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil, e.ID
	}
	return false, r.Error, 0
}

func (a *articleContentDao) Update(e *entity.ArticleContent, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Model(e).Updates(e)
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil
	}
	return false, r.Error
}

func (a *articleContentDao) Delete(e *entity.ArticleContent, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Delete(e)
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil
	}
	return false, r.Error
}

func (a *articleContentDao) SelectOne(id uint64) *entity.ArticleContent {
	var c entity.ArticleContent
	a.db.Where("id = ?", id).Find(&c)
	return &c
}

func (a *articleContentDao) SelectList(conditions ...interface{}) []entity.ArticleContent {
	return make([]entity.ArticleContent, 0)
}

func (a *articleContentDao) SaveSourceItem(source entity.ArticleSource, docId uint64, content string, txContext *TxContext) (bool, error) {
	exist, dbData := a.SelectSourceItem(source, docId)
	update := entity.ArticleContent{
		Content: content,
		DocId:   docId,
		Source:  (uint8)(source),
	}
	var r *gorm.DB
	if exist {
		update.ID = dbData.ID
		r = a.FindDb(txContext, a.db).Model(&update).Update("content", update.Content)
	} else {
		r = a.FindDb(txContext, a.db).Create(&update)
	}
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil
	}
	return false, r.Error
}

func (a *articleContentDao) SelectSourceItem(source entity.ArticleSource, docId uint64) (bool, *entity.ArticleContent) {
	data := entity.ArticleContent{}
	a.db.Where("source = ? and doc_id = ?", source, docId).Find(&data)
	if data.ID > 0 {
		return true, &data
	} else {
		return false, nil
	}
}

func (a *articleContentDao) DeleteSourceItem(source entity.ArticleSource, docId uint64, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Where("source = ? and doc_id = ?", source, docId).Delete(&entity.ArticleContent{})
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil
	}
	return false, r.Error
}
