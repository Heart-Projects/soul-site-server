package entity

import "time"

type ArticleTagRelation struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	// 文章ID
	ArticleId uint64
	// 标签ID
	TagId uint64
}
