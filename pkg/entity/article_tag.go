package entity

type ArticleTag struct {
	BaseModel
	// 用户id
	UserId uint64 `json:"userId"`
	// 标签名称
	Name string `json:"name"`
	// 说明
	Note string `json:"note"`
	// 颜色
	Color string `json:"color"`
}

type ArticleTagPo struct {
	// 用户id
	UserId uint64
	// 标签名称
	Name string `form:"name" json:"name" binding:"required"`
	// 说明
	Note string
	// 颜色
	Color string
}
