package entity

// Article 文章类型实体
type Article struct {
	BaseModel
	// 用户id
	UserId         uint64 `form:"userId" json:"userId"`
	SiteCategoryId uint64 `form:"siteCategoryId" json:"siteCategoryId"`
	// 所属类别
	CategoryId uint64 `form:"categoryId" json:"categoryId"`
	// 专栏id
	ColumnId uint64 `form:"columnId" json:"columnId"`
	// 类型: 0-文章分类 1-普通文章，2-专栏文章
	Type int8 `form:"type" json:"type"`
	// 类别层次路径
	CategoryTreePath string `form:"categoryTreePath" json:"categoryTreePath"`
	// 专栏层次路径
	ColumnTreePath string `form:"columnTreePath" json:"columnTreePath"`
	// 标题/文章分类名称
	Title string `form:"title" json:"title"`
	// 摘要
	Summary string `form:"summary" json:"summary"`
	// 缩略图
	Thumbnail string `form:"thumbnail" json:"thumbnail"`
	// 内容
	Content string `form:"content" json:"content"`
	// 访问量
	ViewCount int `json:"viewCount"`
	// 评论数
	CommentCount int `json:"commentCount"`
	// 点赞数
	LikeCount int `json:"likeCount"`
	// 是否允许评论
	IsComment int8 `form:"isComment" json:"isComment"`
	// 是否置顶
	IsOnTop int8 `form:"isOnTop" json:"isOnTop"`
	// 状态： 0- 草稿 1-发布 2- 垃圾箱
	Status int8 `form:"status" json:"status"`
	// 排序值
	Order int8 `form:"order" json:"order"`
}

// ArticleContent 文章内容
type ArticleContent struct {
	BaseModel
	// Source 内容来源 见 ArticleStatus
	Source uint8 `json:"source"`
	// 文章id
	DocId uint64 `json:"docId"`
	// 文章内容
	Content string `json:"content"`
}

type ArticlePo struct {
	ID uint64 `form:"id" json:"id"`
	// 用户id
	UserId uint64 `json:"userId"`
	// 所属站点类别
	SiteCategoryId uint64 `form:"siteCategoryId" json:"siteCategoryId"`
	// 所属类别
	CategoryId uint64 `form:"categoryId" json:"categoryId"`
	// 专栏id
	ColumnId uint64 `form:"columnId" json:"columnId"`
	// 类型: 0-文章分类 1-普通文章，2-专栏文章
	Type int8 `form:"type" json:"type"`
	// 标题/文章分类名称
	Title string `form:"title" json:"title" binding:"required"`
	// 摘要
	Summary string `form:"summary" json:"summary"`
	// 缩略图
	Thumbnail string `json:"thumbnail"`
	// 内容
	Content string `form:"content" json:"content" binding:"required"`
	// 标签
	Label string `form:"label" json:"label"`
	// 是否允许评论
	IsComment int8 `form:"isComment" json:"isComment"`
	// 是否置顶
	IsOnTop int8 `form:"isOnTop" json:"isOnTop"`
	// 状态： 0- 草稿 1-发布 2- 垃圾箱
	Status int8 `form:"status" json:"status" binding:"required,gte=0,lte=2"`
	// 排序值
	Order int8 `json:"order"`
}

// ArticleType 文章类型
type ArticleType int8

const (
	_ ArticleType = iota
	TypeNormal
	TypeColumn
)

type ArticleSource uint8

const (
	// ArticleSourceDefault 文章来源: 0- 普通文章
	ArticleSourceDefault ArticleSource = iota
	// ArticleSourceColumn 文章来源: 1- 专栏文章
	ArticleSourceColumn
)

// ArticleStatus 文章状态
type ArticleStatus int8

const (
	ArticleStatusDraft ArticleStatus = iota
	ArticleStatusPublish
	ArticleStatusCrash
)

// CommentStatus 是否允许评论
type CommentStatus int8

const (
	ForbiddenComment CommentStatus = iota
	EnableComment
)

// TopStatus 置顶常量
type TopStatus int8

const (
	DisableOnTop TopStatus = iota
	EnableOnTop
)
