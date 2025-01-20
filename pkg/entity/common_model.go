package entity

import (
	"com.sj/admin/pkg/types"
)

type BaseModel struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt types.DateTime `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt types.DateTime `json:"updatedAt " gorm:"autoUpdateTime"`
}
