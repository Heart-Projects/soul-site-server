package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type IColumnArticlesService interface {
	New(vo *vo.ColumnArticlesAddVo) (bool, error, *entity.ColumnArticles)

	Rename(newName string, id uint64) (bool, error)

	ExchangeOrder(sourceItemId uint64, targetItemId uint64) (bool, error)

	ModifyParent(id uint64, parentId uint64) (bool, error)

	SaveContent(id uint64, content string) (bool, error)

	ContentItem(params *vo.ColumnArticleContentVo) (bool, error, *vo.ColumnArticleVo)

	Remove(id uint64) (bool, error)

	GetOne(id uint64) (bool, *entity.ColumnArticles)

	GetByIdentify(identify string) (bool, *entity.ColumnArticles)

	GetColumnData(columnId uint64) (bool, []entity.ColumnArticles)
}

type columnArticlesService struct {
	columnArticlesDao dao.IColumnArticlesDao
	articleContentDao dao.IArticleContentDao
}

func NewColumnArticlesService(columnArticlesDao dao.IColumnArticlesDao, articleContentDao dao.IArticleContentDao) IColumnArticlesService {
	return &columnArticlesService{
		columnArticlesDao: columnArticlesDao,
		articleContentDao: articleContentDao,
	}
}

func (a *columnArticlesService) New(vo *vo.ColumnArticlesAddVo) (bool, error, *entity.ColumnArticles) {
	var data entity.ColumnArticles
	copier.Copy(&data, vo)
	identify, err := utils.RandomString(10)
	if err != nil {
		logrus.Error("save column articles error: ", err)
		return false, err, nil
	}
	data.Identify = identify
	newDocId := uint64(0)
	// 查询当前层级下最大的排序
	currentOrder := a.columnArticlesDao.SelectColumnUnderParentMaxOrder(vo.ColumnId, vo.ParentId)
	data.SortNumber = currentOrder + 1
	err = dao.NewTransactionManager().Start(func(ctx *dao.TxContext) error {
		success, err, newId := a.columnArticlesDao.Insert(&data, nil)
		if !success {
			return err
		}
		contentSuccess, err := a.articleContentDao.SaveSourceItem(entity.ArticleSourceColumn, newId, "", ctx)
		if !contentSuccess {
			return err
		}
		newDocId = newId
		return nil
	})
	if err != nil {
		return false, err, nil
	}
	if querySuccess, savedData := a.GetOne(newDocId); querySuccess {
		return true, nil, savedData
	}
	return false, errors.New("未查询到保存的数据"), nil
}

func (a *columnArticlesService) Rename(newName string, id uint64) (bool, error) {
	return a.columnArticlesDao.Update(&entity.ColumnArticles{
		BaseModel: entity.BaseModel{
			ID: id,
		},
		Title: newName,
	}, nil)
}

// ExchangeOrder 交换排序，排查采用降序排列，即同一层级中值越大的，排序越靠前
func (a *columnArticlesService) ExchangeOrder(sourceItemId uint64, targetItemId uint64) (bool, error) {
	targetArticle := a.columnArticlesDao.SelectOne(targetItemId)
	if targetArticle == nil {
		logrus.Error("target article %d not exist: ", targetItemId)
		return false, errors.New(fmt.Sprintf("专栏文章: %d 不存在", targetItemId))
	}

	sourceArticle := a.columnArticlesDao.SelectOne(sourceItemId)
	if sourceArticle == nil {
		logrus.Error("source article %d not exist: ", sourceItemId)
		return false, errors.New(fmt.Sprintf("专栏文章: %d 不存在", sourceItemId))
	}
	// source 总是插入到target 后面，因此排序好
	newOrder := targetArticle.SortNumber - 1
	updater := &entity.ColumnArticles{
		BaseModel: entity.BaseModel{
			ID: sourceArticle.ID,
		},
		ParentId:   targetArticle.ParentId,
		SortNumber: targetArticle.SortNumber - 1,
	}
	nearestItem := a.columnArticlesDao.SelectFirstAfterOrder(targetArticle.ColumnId, targetArticle.ParentId, newOrder)
	var nearestOrder = 0
	if nearestItem != nil {
		nearestOrder = nearestItem.SortNumber
	}
	// 说明新查询的排序和后续的的最大排序值之间有 未使用的排序值，则直接插入就好
	if nearestOrder < newOrder {
		return a.columnArticlesDao.Update(updater, nil)
	} else {
		// 先更新目标位置只有的数据的排序
		err := dao.NewTransactionManager().Start(func(ctx *dao.TxContext) error {
			ok, err := a.columnArticlesDao.UpdateWithDecreaseOrderAfter(targetArticle.ColumnId, targetArticle.ParentId, newOrder, ctx)
			if !ok {
				return err
			}
			ok, err = a.columnArticlesDao.Update(updater, ctx)
			if !ok {
				return err
			}
			return nil
		})
		return err == nil, errors.New("更新专栏文章排序失败")
	}

}

func (a *columnArticlesService) ModifyParent(id uint64, parentId uint64) (bool, error) {
	return a.columnArticlesDao.Update(&entity.ColumnArticles{
		BaseModel: entity.BaseModel{
			ID: id,
		},
		ParentId: parentId,
	}, nil)
}

func (a *columnArticlesService) SaveContent(id uint64, content string) (bool, error) {
	return a.articleContentDao.SaveSourceItem(entity.ArticleSourceColumn, id, content, nil)
}

func (a *columnArticlesService) ContentItem(params *vo.ColumnArticleContentVo) (bool, error, *vo.ColumnArticleVo) {
	_, columnArticle := a.columnArticlesDao.SelectByIdentify(params.Identify)

	if columnArticle == nil {
		return false, errors.New("专栏文章不存在"), nil

	}
	success, articleContent := a.articleContentDao.SelectSourceItem(entity.ArticleSourceColumn, columnArticle.ID)
	if !success {
		return false, errors.New("专栏文章不存在"), nil
	}
	return true, nil, &vo.ColumnArticleVo{
		ColumnArticles: columnArticle,
		Content:        articleContent.Content,
	}
}

// Remove 删除专栏文章
func (a *columnArticlesService) Remove(id uint64) (bool, error) {
	err := dao.NewTransactionManager().Start(func(ctx *dao.TxContext) error {
		success, err := a.columnArticlesDao.Delete(&entity.ColumnArticles{
			BaseModel: entity.BaseModel{
				ID: id,
			},
		}, ctx)
		if !success {
			return err
		}
		contentSuccess, contentErr := a.articleContentDao.DeleteSourceItem(entity.ArticleSourceColumn, id, ctx)
		if !contentSuccess {
			return contentErr
		}
		return nil
	})
	return err == nil, err
}

func (a *columnArticlesService) GetOne(id uint64) (bool, *entity.ColumnArticles) {
	data := a.columnArticlesDao.SelectOne(id)
	return data != nil, data
}

func (a *columnArticlesService) GetByIdentify(identify string) (bool, *entity.ColumnArticles) {
	return a.columnArticlesDao.SelectByIdentify(identify)
}

func (a *columnArticlesService) GetColumnData(columnId uint64) (bool, []entity.ColumnArticles) {
	return true, a.columnArticlesDao.SelectList(columnId)
}

func (a *columnArticlesService) GetColumnUnderParentMaxOrder(columnId uint64, parentId uint64) int {
	return a.columnArticlesDao.SelectColumnUnderParentMaxOrder(columnId, parentId)
}
