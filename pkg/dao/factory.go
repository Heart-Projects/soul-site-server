package dao

import (
	"com.sj/admin/pkg/factory"
	"gorm.io/gorm"
)

type Factory interface {
	Article() IArticleDao
	ArticleCategory() IArticleCategoryDao
	User() IUserDao
	SiteCategory() ISiteCategoryDao

	ArticleColumnCategory() IArticleColumnCategoryDao

	ArticleColumn() IArticleColumnDao

	ArticleContent() IArticleContentDao

	ColumnArticles() IColumnArticlesDao

	GetDb() *gorm.DB
}

type mysqlFactory struct {
	db *gorm.DB
}

func (m *mysqlFactory) Article() IArticleDao {
	return factory.GetOrCreateIns[IArticleDao](func(params ...any) interface{} {
		return newArticleDao(m.db)
	})
}

func (m *mysqlFactory) ArticleCategory() IArticleCategoryDao {
	return factory.GetOrCreateIns[IArticleCategoryDao](func(params ...any) interface{} {
		return newArticleCategoryDao(m.db)
	})
}

func (m *mysqlFactory) User() IUserDao {
	return factory.GetOrCreateIns[IUserDao](func(params ...any) interface{} {
		return newUserDao(m.db)
	})
}

func (m *mysqlFactory) SiteCategory() ISiteCategoryDao {
	return factory.GetOrCreateIns[ISiteCategoryDao](func(params ...interface{}) interface{} {
		return newSiteCategoryDao(m.db)
	})
}

func (m *mysqlFactory) ArticleColumnCategory() IArticleColumnCategoryDao {
	return factory.GetOrCreateIns[IArticleColumnCategoryDao](func(params ...any) interface{} {
		return newArticleColumnCategoryDao(m.db)
	})
}

func (m *mysqlFactory) ArticleColumn() IArticleColumnDao {
	return factory.GetOrCreateIns[IArticleColumnDao](func(params ...any) interface{} {
		return newArticleColumnDao(m.db)
	})
}

func (m *mysqlFactory) ArticleContent() IArticleContentDao {
	return factory.GetOrCreateIns[IArticleContentDao](func(params ...any) interface{} {
		return newArticleContentDao(m.db)
	})
}

func (m *mysqlFactory) ColumnArticles() IColumnArticlesDao {
	return factory.GetOrCreateIns[IColumnArticlesDao](func(params ...any) interface{} {
		return newColumnArticlesDao(m.db)
	})
}

func (m *mysqlFactory) GetDb() *gorm.DB {
	return m.db
}

func NewOrGet() Factory {
	return factory.GetOrCreateIns[Factory](func(params ...any) interface{} {
		db := factory.GetDbInstance()
		return &mysqlFactory{db: db}
	})

}
