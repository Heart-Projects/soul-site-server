package entity

type ArticleStatistics struct {
	BaseModel
	UserId    uint64 `json:"userId"`
	ArticleId uint64 `json:"articleId"`
	// 文章阅读数
	ViewCount int
	// 文章点赞数
	LikeCount int
	// 文章评论数
	CommentCount int
	// 文章转发数
	ForwardCount int
	// 文章收藏数
	FavoriteCount int
}
