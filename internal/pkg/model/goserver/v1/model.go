package v1

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel is to replace `gorm.BaseModel`.
type BaseModel struct {
	ID        uint64         `json:"id,omitempty"        gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt;index:idx_deletedAt"`
}
