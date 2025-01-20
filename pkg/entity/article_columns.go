package entity

type ArticleColumnCategory struct {
	BaseModel
	// ParentId 父级ID
	ParentId uint64 `json:"parentId"`
	// UserId 用户ID
	UserId uint64 `json:"userId"`
	Name   string `json:"name"`
	// Level 级别
	Level uint8 `json:"level"`
	// 树路径
	TreePath string `json:"treePath"`
}

// ArticleColumnTheme 专栏主题
type ArticleColumnTheme struct {
	BaseModel
	Name string `json:"name"`
	Note string `json:"note"`
	Icon string `json:"icon"`
}

// ArticleColumn 专栏
type ArticleColumn struct {
	BaseModel
	// CategoryId 专栏类别ID
	CategoryId uint64 `json:"categoryId"`
	// ThemeId 专栏主题ID
	ThemeId uint64 `json:"themeId"`
	// UserId 用户ID
	UserId   uint64 `json:"userId"`
	Identify string `json:"identify"`
	Name     string `json:"name"`
	Note     string `json:"note"`
	Icon     string `json:"icon"`
	Summary  string `json:"summary"`
	// ViewScope 查看权限
	ViewScope int `json:"viewScope"`
}

// ColumnArticles 专栏文章
type ColumnArticles struct {
	BaseModel
	ParentId   uint64 `json:"parentId"`
	ColumnId   uint64 `json:"columnId"`
	Type       string `json:"type"`
	TreePath   string `json:"treePath"`
	Identify   string `json:"identify"`
	Title      string `json:"title"`
	SortNumber int    `json:"sortNumber"`
}

const (
	ColumnArticleTypeDoc    string = "doc"
	ColumnArticleTypeFolder string = "folder"
	ColumnArticleTypeExcel  string = "excel"
)

// IsLegalColumnArticleType 是否合法的专栏文章类型
func IsLegalColumnArticleType(t string) bool {
	switch t {
	case ColumnArticleTypeDoc, ColumnArticleTypeFolder, ColumnArticleTypeExcel:
		return true
	}
	return false

}
