package vo

import "com.sj/admin/pkg/entity"

type UserArticlePageVo struct {
	Total     int64           `json:"total"`
	List      []UserArticleVo `json:"list"`
	UserId    uint64          `json:"userId"`
	PageIndex int             `form:"pageIndex" json:"pageIndex" binding:"required"`
	PageSize  int             `form:"pageSize" json:"pageSize" binding:"required"`
	Sort      string          `json:"sort"`
}

type UserArticleVo struct {
	entity.Article
	// 文章标签
	Labels *[]entity.ArticleTag `json:"labels"`
	// 下一遍文章
	Next *UserNavArticleVo `json:"next"`
	// 上一遍文章
	Pre *UserNavArticleVo `json:"pre"`
}

type ArticleDetailParams struct {
	ArticleId uint64 `uri:"articleId" json:"articleId" binding:"required"`

	UserId uint64 `uri:"userId" json:"userId" binding:"required"`
}

type UserNavArticleVo struct {
	ID uint64 `json:"id"`

	Title string `json:"title"`
}
