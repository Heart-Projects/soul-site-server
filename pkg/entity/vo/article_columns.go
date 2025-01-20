package vo

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/utils"
)

// ArticleColumnCategoryAddVo 新增
type ArticleColumnCategoryAddVo struct {
	// ParentId 父级ID
	ParentId uint64 `form:"parentId" json:"parentId"`
	// UserId 用户ID
	UserId uint64 `json:"userId"`
	Name   string `form:"name" binding:"required" json:"name"`
}

// ArticleColumnCategoryUpdateVo 更新
type ArticleColumnCategoryUpdateVo struct {
	ID uint64 `form:"id" binding:"required" json:"id"`
	// ParentId 父级ID
	ParentId uint64 `form:"parentId" json:"parentId"`

	Name string `form:"name" binding:"required" json:"name"`
}

type ArticleColumnAddVo struct {
	CategoryId uint64 `form:"categoryId" binding:"required" json:"categoryId"`
	// ThemeId 专栏主题ID
	ThemeId uint64 `form:"themeId" json:"themeId"`
	// UserId 用户ID
	UserId   uint64 `json:"userId"`
	Identify string `json:"identify"`
	Name     string `form:"name" binding:"required" json:"name"`
	Note     string `form:"note" json:"note"`
	Icon     string `form:"icon" json:"icon"`
	Summary  string `form:"summary" json:"summary"`
	// ViewScope 查看权限
	ViewScope int ` form:"viewScope" json:"viewScope"`
}

type ArticleColumnUpdateVo struct {
	ID         uint64 `form:"id" binding:"required" json:"id"`
	CategoryId uint64 `form:"categoryId" binding:"required" json:"categoryId"`
	// ThemeId 专栏主题ID
	ThemeId uint64 `form:"themeId" json:"themeId"`
	Name    string `form:"name" binding:"required" json:"name"`
	Note    string `form:"note" json:"note"`
	Icon    string `form:"icon" json:"icon"`
	// ViewScope 查看权限
	ViewScope int ` form:"viewScope" json:"viewScope"`
}

type ArticleColumnCategoryListVo struct {
	*entity.ArticleColumnCategory
	Children []utils.TreeItemFeature `json:"children"`
}

func (a *ArticleColumnCategoryListVo) GetId() uint64 {
	return a.ID
}

func (a *ArticleColumnCategoryListVo) GetParentId() uint64 {
	return a.ParentId
}

func (a *ArticleColumnCategoryListVo) SetChildren(children []utils.TreeItemFeature) {
	if a.Children == nil {
		a.Children = make([]utils.TreeItemFeature, 0)
	}
	a.Children = children
}

func (a *ArticleColumnCategoryListVo) GetChildren() []utils.TreeItemFeature {
	return a.Children
}

type ColumnArticlesAddVo struct {
	ParentId       uint64 `form:"parentId" json:"parentId"`
	ColumnId       uint64 `form:"columnId" json:"columnId"`
	ColumnIdentify string `form:"columnIdentify" json:"columnIdentify"`
	Type           string `form:"type" binding:"required" json:"type"`
	Title          string `form:"title" binding:"required" json:"title"`
}

type ColumnArticleRenameVo struct {
	Id    uint64 `form:"id" binding:"required" json:"id"`
	Title string `form:"title" binding:"required" json:"title"`
}

type ColumnArticleSaveContent struct {
	Id      uint64 `form:"id" binding:"required" json:"id"`
	Content string `form:"content" binding:"required" json:"content"`
}

type ColumnArticleModifyOrderVo struct {
	// 被移动的项目ID
	SourceItemId uint64 `form:"sourceItemId" binding:"required" json:"sourceItemId"`
	// 目标项目ID
	TargetItemId uint64 `form:"targetItemId" binding:"required" json:"targetItemId"`
}

type ColumnArticleModifyParentVo struct {
	Id       uint64 `form:"id" binding:"required" json:"id"`
	ParentId uint64 `form:"parentId" binding:"required" json:"parentId"`
}

type ColumnArticleVo struct {
	*entity.ColumnArticles
	Content interface{} `json:"content"`
}

type ColumnArticleContentVo struct {
	ColumnIdentify string `form:"columnIdentify" binding:"required" json:"columnIdentify"`
	Identify       string `form:"identify"  json:"identify"`
}

type ColumnArticlesListVo struct {
	*entity.ColumnArticles
	Children []utils.TreeItemFeature `json:"children"`
}

func (a *ColumnArticlesListVo) GetId() uint64 {
	return a.ID
}

func (a *ColumnArticlesListVo) GetParentId() uint64 {
	return a.ParentId
}

func (a *ColumnArticlesListVo) SetChildren(children []utils.TreeItemFeature) {
	if a.Children == nil {
		a.Children = make([]utils.TreeItemFeature, 0)
	}
	a.Children = children
}

func (a *ColumnArticlesListVo) GetChildren() []utils.TreeItemFeature {
	return a.Children
}

type ArticleContentRequestVo struct {
	Source      uint8  `form:"source" binding:"required" json:"source"`
	DocIdentify string `form:"docIdentify" json:"docIdentify"`
	DocId       uint64 `form:"docId"  json:"docId"`
}

type ArticleContentSaveVo struct {
	Source  uint8  `form:"source" binding:"required" json:"source"`
	DocId   uint64 `form:"docId" binding:"required" json:"docId"`
	Content string `form:"content" binding:"required" json:"content"`
}
