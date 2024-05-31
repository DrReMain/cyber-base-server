package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID        uint64                `gorm:"primaryKey;autoIncrement;"`
	CreatedAt uint64                `gorm:"autoCreateTime:milli;comment:创建时间"`
	UpdatedAt uint64                `gorm:"autoUpdateTime:milli;comment:更新时间"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;index:idx_mobile_deleted_at;comment:删除时间"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = uint64(time.Now().UnixMilli())
	return nil
}

func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = uint64(time.Now().UnixMilli())
	return nil
}
