package dao

import (
	"com.sj/admin/pkg/entity"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

// IColumnArticlesDao 专栏文章
type IColumnArticlesDao interface {
	IBaseDao[entity.ColumnArticles]
	SelectByIdentify(identify string) (bool, *entity.ColumnArticles)

	SelectColumnUnderParentMaxOrder(columnId uint64, parentId uint64) int
	// SelectFirstAfterOrder 查询查询当前层级下的最近的元素
	SelectFirstAfterOrder(columnId uint64, parentId uint64, targetOrder int) *entity.ColumnArticles
	// UpdateWithDecreaseOrderAfter 更新当前排序后的元素
	UpdateWithDecreaseOrderAfter(columnId uint64, parentId uint64, targetOrder int, txContext *TxContext) (bool, error)
}

type columnArticlesDao struct {
	db *gorm.DB
	BaseDao[entity.ColumnArticles]
}

func newColumnArticlesDao(db *gorm.DB) IColumnArticlesDao {
	return &columnArticlesDao{db: db}
}

func (a *columnArticlesDao) Insert(e *entity.ColumnArticles, txContext *TxContext) (bool, error, uint64) {
	if success, err := a.fillExtraInfo(e.ParentId, e); !success {
		return false, err, 0
	}
	r := a.FindDb(txContext, a.db).Create(e)
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil, e.ID
	}
	return false, r.Error, 0
}

func (a *columnArticlesDao) Update(e *entity.ColumnArticles, txContext *TxContext) (bool, error) {
	if success, err := a.fillExtraInfo(e.ParentId, e); !success {
		return false, err
	}
	r := a.FindDb(txContext, a.db).Model(e).Updates(e)
	if r.Error == nil && r.RowsAffected == 1 {
		return true, nil
	}
	logrus.Error("update error: ", r.Error)
	return false, r.Error
}

func (a *columnArticlesDao) Delete(e *entity.ColumnArticles, txContext *TxContext) (bool, error) {
	if txContext == nil {
		err := NewTransactionManager().Start(func(ctx *TxContext) error {
			return a.doDelete(e, ctx)
		})
		return err == nil, err
	}
	err := a.doDelete(e, txContext)
	return err == nil, a.doDelete(e, txContext)
}

// doDelete 删除当前节点及其子几点
func (a *columnArticlesDao) doDelete(e *entity.ColumnArticles, txContext *TxContext) error {
	r := a.FindDb(txContext, a.db).Delete(e)
	if r.Error != nil || r.RowsAffected == 0 {
		logrus.Errorf("Column Article not exist: %d", e.ID)
		return errors.New(fmt.Sprintf("Column Article not exist: %d", e.ID))
	}
	// 根据treePath删除子节点
	r = a.FindDb(txContext, a.db).Where(" tree_path like ? ", "/"+strconv.FormatUint(e.ID, 10)+"/%").Delete(&entity.ColumnArticles{})
	if r.Error != nil {
		logrus.Error("delete column article children failed, error: ", r.Error)
		return r.Error
	}
	return nil
}

func (a *columnArticlesDao) SelectOne(id uint64) *entity.ColumnArticles {
	var c entity.ColumnArticles
	a.db.Where("id = ?", id).Find(&c)
	return &c
}

func (a *columnArticlesDao) SelectList(conditions ...interface{}) []entity.ColumnArticles {
	var c []entity.ColumnArticles
	a.db.Where(" column_id = ?", conditions).Find(&c)
	return c
}

func (a *columnArticlesDao) SelectByIdentify(identify string) (bool, *entity.ColumnArticles) {
	var c entity.ColumnArticles
	a.db.Where("identify = ?", identify).Find(&c)
	if c.ID == 0 {
		return false, nil
	}
	return true, &c
}

func (a *columnArticlesDao) SelectColumnUnderParentMaxOrder(columnId uint64, parentId uint64) int {
	var c entity.ColumnArticles
	a.db.Where("column_id = ? and parent_id = ?", columnId, parentId).Order("sort_number desc").Find(&c)
	if c.ID == 0 {
		return 0
	}
	return c.SortNumber
}

func (a *columnArticlesDao) SelectFirstAfterOrder(columnId uint64, parentId uint64, targetOrder int) *entity.ColumnArticles {
	var c entity.ColumnArticles
	a.db.Where("column_id = ? and parent_id = ? and sort_number <= ?", columnId, parentId, targetOrder).Order("sort_number desc").Limit(1).Find(&c)
	return &c
}

func (a *columnArticlesDao) UpdateWithDecreaseOrderAfter(columnId uint64, parentId uint64, targetOrder int, txContext *TxContext) (bool, error) {
	r := a.FindDb(txContext, a.db).Model(&entity.ColumnArticles{}).Where("column_id = ? and parent_id = ? and sort_number < ?", columnId, parentId, targetOrder).Update("sort_number", gorm.Expr("sort_number - ?", 1))
	if r.Error == nil {
		return true, nil
	}
	return false, r.Error
}

// fillExtraInfo 根据父节点填充 treePath 和 level
func (a *columnArticlesDao) fillExtraInfo(parentId uint64, e *entity.ColumnArticles) (bool, error) {
	if parentId == 0 {
		e.TreePath = "/"
		return true, nil
	}
	p := a.SelectOne(e.ParentId)
	if p == nil {
		return false, errors.New("parent not exist")
	}
	e.TreePath = p.TreePath + strconv.FormatUint(parentId, 10) + "/"
	return true, nil
}
