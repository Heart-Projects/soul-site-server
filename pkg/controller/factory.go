package controller

import (
	"com.sj/admin/pkg/dao"
	"com.sj/admin/pkg/factory"
	"com.sj/admin/pkg/service"
)

type ControllerFactory interface {
	UserController() IUserController
	TagController() ITagController
	ArticleController() IArticleController
	ArticleCategoryController() IArticleCategoryController
	SiteCategoryController() ISiteCategoryController

	ArticleColumnCategoryController() IArticleColumnCategoryController

	ArticleColumnController() IArticleColumnController

	ColumnArticlesController() IColumnArticlesController

	ArticleContentController() IArticleContentController

	FileController() IFileController
}

type ControllerFactoryImpl struct {
}

// 类型检查
var _ ControllerFactory = (*ControllerFactoryImpl)(nil)

func (c *ControllerFactoryImpl) UserController() IUserController {
	return factory.GetOrCreateIns[IUserController](func(params ...interface{}) interface{} {
		return NewUserController(service.GetOrCreate().User())
	})
}

func (c *ControllerFactoryImpl) TagController() ITagController {
	return factory.GetOrCreateIns[ITagController](func(params ...interface{}) interface{} {
		return NewTagController(service.GetOrCreate().Tag())
	})
}

func (c *ControllerFactoryImpl) ArticleController() IArticleController {
	return factory.GetOrCreateIns[IArticleController](func(params ...interface{}) interface{} {
		return NewArticleController(service.GetOrCreate().Article())
	})
}

func (c *ControllerFactoryImpl) ArticleCategoryController() IArticleCategoryController {
	return factory.GetOrCreateIns[IArticleCategoryController](func(params ...interface{}) interface{} {
		return NewArticleCategoryController(service.GetOrCreate().ArticleCategory())
	})
}

func (c *ControllerFactoryImpl) SiteCategoryController() ISiteCategoryController {
	return factory.GetOrCreateIns[ISiteCategoryController](func(params ...interface{}) interface{} {
		return NewSiteCategoryController(dao.NewOrGet().SiteCategory())
	})
}

func (c *ControllerFactoryImpl) ArticleColumnCategoryController() IArticleColumnCategoryController {
	return factory.GetOrCreateIns[IArticleColumnCategoryController](func(params ...interface{}) interface{} {
		return NewArticleColumnCategoryController(service.GetOrCreate().ArticleColumnCategory())
	})
}

func (c *ControllerFactoryImpl) ArticleColumnController() IArticleColumnController {
	return factory.GetOrCreateIns[IArticleColumnController](func(params ...interface{}) interface{} {
		return NewArticleColumnController(service.GetOrCreate().ArticleColumn())
	})
}

func (c *ControllerFactoryImpl) ColumnArticlesController() IColumnArticlesController {
	return factory.GetOrCreateIns[IColumnArticlesController](func(params ...interface{}) interface{} {
		return NewColumnArticlesController(service.GetOrCreate().ColumnArticles(), service.GetOrCreate().ArticleColumn())
	})
}

func (c *ControllerFactoryImpl) ArticleContentController() IArticleContentController {
	return factory.GetOrCreateIns[IArticleContentController](func(params ...interface{}) interface{} {
		return NewArticleContentController(service.GetOrCreate().ArticleContent())
	})
}

func (c *ControllerFactoryImpl) FileController() IFileController {
	return factory.GetOrCreateIns[IFileController](func(params ...interface{}) interface{} {
		return NewFileController()
	})
}

func getOrCreate() ControllerFactory {
	return factory.GetOrCreateIns[ControllerFactory](func(params ...interface{}) interface{} {
		return &ControllerFactoryImpl{}
	})
}
