package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/utils"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type IArticleColumnService interface {
	New(param *vo.ArticleColumnAddVo) (bool, error, *entity.ArticleColumn)
	Modify(param *vo.ArticleColumnUpdateVo) (bool, error)
	Remove(id uint64) (bool, error)
	// GetOne 获取指定的文章专栏，支持ID 和 identify
	GetOne(id uint64, identify ...string) *entity.ArticleColumn
	List(categoryId uint64) []entity.ArticleColumn
}

type articleColumnService struct {
	columnDao dao.IArticleColumnDao
}

func NewArticleColumnService(columnDao dao.IArticleColumnDao) IArticleColumnService {
	return &articleColumnService{
		columnDao: columnDao,
	}
}

func (a *articleColumnService) New(param *vo.ArticleColumnAddVo) (bool, error, *entity.ArticleColumn) {
	data := &entity.ArticleColumn{}
	err := copier.Copy(data, param)
	if err != nil {
		return false, err, nil
	}
	identify, err := utils.RandomString(12)
	if err != nil {
		logrus.Error("生成唯一标识失败", err)
		return false, errors.New("生成唯一标识失败"), nil
	}
	data.Identify = identify
	success, err, newId := a.columnDao.Insert(data, nil)
	if success {
		data.ID = newId
		return true, nil, data
	}
	return false, err, nil
}

func (a *articleColumnService) Modify(param *vo.ArticleColumnUpdateVo) (bool, error) {
	var updateData = &entity.ArticleColumn{}
	if err := copier.Copy(updateData, param); err != nil {
		logrus.WithError(err).Error("拷贝数据失败")
		return false, err
	}
	return a.columnDao.Update(updateData, nil)
}

func (a *articleColumnService) Remove(id uint64) (bool, error) {
	return a.columnDao.Delete(&entity.ArticleColumn{BaseModel: entity.BaseModel{ID: id}}, nil)
}

func (a *articleColumnService) GetOne(id uint64, identify ...string) *entity.ArticleColumn {
	if id > 0 {
		return a.columnDao.SelectOne(id)
	}
	if len(identify) > 0 {
		return a.columnDao.SelectOneByField("identify", identify[0])
	}
	return nil
}

func (a *articleColumnService) List(categoryId uint64) []entity.ArticleColumn {
	return a.columnDao.SelectList(categoryId)
}
