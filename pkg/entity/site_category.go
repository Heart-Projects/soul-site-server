package entity

type SiteCategory struct {
	BaseModel
	// 父分类id
	ParentId uint64 `json:"parentId"`
	// 类别名称
	Name string `json:"name"`
	// 说明
	Note string `json:"note"`
	// 显示顺序
	Order uint16 `json:"order"`
	// 是否显示：0- 不显示 1- 显示
	IsShow uint8 `json:"isShow"`
}

// 树形结构
type TreeSiteCategory struct {
	SiteCategory
	Children []*TreeSiteCategory `json:"children"`
}
