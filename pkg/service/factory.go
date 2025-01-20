package service

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/factory"
)

// 服务工厂
type ServiceFactory interface {
	ArticleCategory() IArticleCategoryService
	Article() IArticleService
	Tag() ITagService
	ArticleTag() IArticleTagService
	User() IUserService

	ArticleColumnCategory() IArticleColumnCategoryService

	ArticleColumn() IArticleColumnService

	ArticleContent() IArticleContentService

	ColumnArticles() IColumnArticlesService
}

// 服务实现
type serviceFactoryImpl struct {
}

// 通过类型变量赋值语句巧妙的检查 serviceFactoryImpl 必须实现 ServiceFactory
var _ ServiceFactory = (*serviceFactoryImpl)(nil)

func (s serviceFactoryImpl) ArticleCategory() IArticleCategoryService {
	return factory.GetOrCreateIns[IArticleCategoryService](func(params ...any) interface{} {
		return newArticleCategory(dao.NewOrGet().ArticleCategory())
	})
}

func (s serviceFactoryImpl) Article() IArticleService {
	return factory.GetOrCreateIns[IArticleService](func(params ...any) interface{} {
		return NewArticleService(s.ArticleTag(), s.Tag(), dao.NewOrGet().Article())
	})
}

func (s serviceFactoryImpl) Tag() ITagService {
	return factory.GetOrCreateIns[ITagService](func(params ...any) interface{} {
		return NewTagService(s.ArticleTag(), factory.GetDbInstance())
	})
}

func (s serviceFactoryImpl) ArticleTag() IArticleTagService {
	return factory.GetOrCreateIns[IArticleTagService](func(params ...any) interface{} {
		return NewArticleTagService(factory.GetDbInstance())
	})
}

func (s serviceFactoryImpl) User() IUserService {
	return factory.GetOrCreateIns[IUserService](func(params ...any) interface{} {
		return NewUserService(dao.NewOrGet().User())
	})
}

func (s serviceFactoryImpl) ArticleColumnCategory() IArticleColumnCategoryService {
	return factory.GetOrCreateIns[IArticleColumnCategoryService](func(params ...any) interface{} {
		return NewArticleColumnCategoryService(dao.NewOrGet().ArticleColumnCategory())
	})
}

func (s serviceFactoryImpl) ArticleColumn() IArticleColumnService {
	return factory.GetOrCreateIns[IArticleColumnService](func(params ...any) interface{} {
		return NewArticleColumnService(dao.NewOrGet().ArticleColumn())
	})
}

func (s serviceFactoryImpl) ArticleContent() IArticleContentService {
	return factory.GetOrCreateIns[IArticleContentService](func(params ...any) interface{} {
		return NewArticleContentService(dao.NewOrGet().ArticleContent())
	})
}

func (s serviceFactoryImpl) ColumnArticles() IColumnArticlesService {
	return factory.GetOrCreateIns[IColumnArticlesService](func(params ...any) interface{} {
		return NewColumnArticlesService(dao.NewOrGet().ColumnArticles(), dao.NewOrGet().ArticleContent())
	})
}

func GetOrCreate() ServiceFactory {
	return factory.GetOrCreateIns[ServiceFactory](func(params ...interface{}) interface{} {
		return &serviceFactoryImpl{}
	})
}
