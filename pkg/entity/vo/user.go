package vo

import (
	"com.sj/admin/pkg/entity"
)

type SimpleUserVo struct {
	ID uint64 `gorm:"primarykey" json:"id"`
	// 用户标识
	UserIdentify string `json:"userIdentify"`
	// 用户名
	Name string `json:"name"`
	// 邮箱
	Email string `json:"email"`
	// 网站主页地址
	HomeUrl string `json:"homeUrl"`
	// 头像
	Avatar string `json:"avatar"`
}
type UserArticleData struct {
	User        *SimpleUserVo                     `json:"user"`
	ArticleData *entity.UserArticleStatisticsData `json:"articleData"`
}
