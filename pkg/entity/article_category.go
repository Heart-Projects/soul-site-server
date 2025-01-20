package entity

type ArticleCategory struct {
	BaseModel
	// 用户id
	UserId uint64 `json:"userId"`
	// 父分类id
	ParentId uint64 `json:"parentId"`
	// 类别名称
	Name string `json:"name"`
	// 说明
	Note string `json:"note"`
	// 显示顺序
	Order uint16 `json:"order"`
	// 级别
	Level uint8 `json:"level"`
	// 是否显示：0- 不显示 1- 显示
	IsShow uint8 `json:"isShow"`
	// 文章数目
	ArticleCount uint `json:"articleCount"`
}

// 树形结构
type TreeArticleCategory struct {
	ArticleCategory
	Children []*TreeArticleCategory `json:"children"`
}
