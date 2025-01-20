package dao

import (
	"com.sj/admin/pkg/entity"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

// IArticleColumnCategoryDao 专栏类型
type IArticleColumnCategoryDao interface {
	IBaseDao[entity.ArticleColumnCategory]
}

type articleColumnCategoryDao struct {
	db *gorm.DB
	BaseDao[entity.ArticleColumnCategory]
}

func newArticleColumnCategoryDao(db *gorm.DB) IArticleColumnCategoryDao {
	return &articleColumnCategoryDao{db: db}
}

func (a *articleColumnCategoryDao) Insert(e *entity.ArticleColumnCategory, txContext *TxContext) (bool, error, uint64) {
	if success, err := a.fillExtraInfo(e.ParentId, e); !success {
		return false, err, 0
	}
	r := a.FindDb(txContext, a.db).Create(e)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil, e.ID
	}
	return false, r.Error, 0
}

func (a *articleColumnCategoryDao) Update(e *entity.ArticleColumnCategory, txContext *TxContext) (bool, error) {
	dbData := a.SelectOne(e.ID)
	if dbData == nil {
		return false, errors.New(fmt.Sprintf("Category not exist: %d", e.ID))
	}
	var updater = &entity.ArticleColumnCategory{
		Name: e.Name,
		BaseModel: entity.BaseModel{
			ID: e.ID,
		},
	}

	if dbData.ParentId != e.ParentId {
		success, err := a.fillExtraInfo(e.ParentId, updater)
		if !success {
			return false, err
		}
	}
	r := a.FindDb(txContext, a.db).Model(updater).Updates(updater)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil
	}
	return false, r.Error
}

func (a *articleColumnCategoryDao) Delete(e *entity.ArticleColumnCategory, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Delete(e)
	if success := r.Error == nil && r.RowsAffected == 1; success {
		return success, nil
	}
	logrus.Error("delete error: ", r.Error)
	return false, errors.New(fmt.Sprintf("Category not exist: %d", e.ID))
}

func (a *articleColumnCategoryDao) SelectOne(id uint64) *entity.ArticleColumnCategory {
	var c entity.ArticleColumnCategory
	a.db.Where("id = ?", id).Find(&c)
	return &c
}

func (a *articleColumnCategoryDao) SelectList(conditions ...interface{}) []entity.ArticleColumnCategory {
	var c []entity.ArticleColumnCategory
	a.db.Find(&c)
	return c
}

// fillExtraInfo 根据父节点填充 treePath 和 level
func (a *articleColumnCategoryDao) fillExtraInfo(parentId uint64, e *entity.ArticleColumnCategory) (bool, error) {
	if parentId == 0 {
		e.Level = 0
		e.TreePath = "/"
		return true, nil
	}
	p := a.SelectOne(e.ParentId)
	if p == nil {
		return false, errors.New("parent not exist")
	}
	e.Level = p.Level + 1
	e.TreePath = p.TreePath + strconv.FormatUint(parentId, 10) + "/"
	return true, nil
}
