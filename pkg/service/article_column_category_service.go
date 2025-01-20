package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
)

type IArticleColumnCategoryService interface {
	New(param *vo.ArticleColumnCategoryAddVo) (bool, error, *entity.ArticleColumnCategory)
	Rename(param *vo.ArticleColumnCategoryUpdateVo) (bool, error)
	Remove(id uint64) (bool, error)
	List(userId uint64) []entity.ArticleColumnCategory
}

type ArticleColumnCategoryService struct {
	columnCategoryDao dao.IArticleColumnCategoryDao
}

func NewArticleColumnCategoryService(categoryDao dao.IArticleColumnCategoryDao) IArticleColumnCategoryService {
	return &ArticleColumnCategoryService{
		columnCategoryDao: categoryDao,
	}
}

func (a *ArticleColumnCategoryService) New(param *vo.ArticleColumnCategoryAddVo) (bool, error, *entity.ArticleColumnCategory) {
	data := &entity.ArticleColumnCategory{
		ParentId: param.ParentId,
		UserId:   param.UserId,
		Name:     param.Name,
	}
	success, err, newId := a.columnCategoryDao.Insert(data, nil)
	if success {
		data.ID = newId
		return true, nil, data
	}
	return false, err, nil
}

func (a *ArticleColumnCategoryService) Rename(param *vo.ArticleColumnCategoryUpdateVo) (bool, error) {
	var updateData = &entity.ArticleColumnCategory{
		ParentId: param.ParentId,
		Name:     param.Name,
		BaseModel: entity.BaseModel{
			ID: param.ID,
		},
	}
	return a.columnCategoryDao.Update(updateData, nil)
}

func (a *ArticleColumnCategoryService) Remove(id uint64) (bool, error) {
	return a.columnCategoryDao.Delete(&entity.ArticleColumnCategory{BaseModel: entity.BaseModel{ID: id}}, nil)
}

func (a *ArticleColumnCategoryService) List(userId uint64) []entity.ArticleColumnCategory {
	params := make(map[string]interface{})
	params["userId"] = userId
	return a.columnCategoryDao.SelectList(params)
}
