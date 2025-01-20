package entity

import (
	"time"
)

type SysUser struct {
	// 注意这里没有使用Model 是由于和jwt中的类型不匹配
	ID        uint64    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// 用户标识
	UserIdentify string
	// 用户名
	Name string
	// 邮箱
	Email string
	// 手机号
	Phone string
	// 网站主页地址
	HomeUrl string
	// 邀请码
	InviteCode string
	// 头像
	Avatar string
	// 是否允许后台登陆
	BackgroundLogin int8
	// 最后一次登陆IP
	LastLoginIp string
	// 最后一次登陆时间
	LastLoginAt time.Time
	// 密码
	Password string
	// 备注
	Remark string
	// 状态
	Status int
}

const (
	// UserBackgroundLoginNo 0 不允许后台登陆
	UserBackgroundLoginNo = iota
	// UserBackgroundLoginYes 0 允许后台登陆
	UserBackgroundLoginYes
)

// UserArticleStatisticsData 用户文章统计数据
type UserArticleStatisticsData struct {
	BaseModel
	// 人气值
	HotCount uint `json:"hotCount"`
	// 文章数
	ArticleCount uint `json:"articleCount"`
	// 关注数
	FollowCount uint `json:"followCount"`
	// 收藏数
	FavoriteCount uint `json:"favoriteCount"`
	// 转发数
	ForwardCount uint `json:"forwardCount"`
	// 评论数
	CommentCount uint `json:"commentCount"`
}
