package models

import (
	"github.com/zqddong/learnku-blog/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;nou null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return types.Uint64ToString(b.ID)
}
