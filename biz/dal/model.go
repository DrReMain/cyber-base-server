package dal

import (
	"log"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("[DB]: snowflake.NewNode -> '%s'\n", err)
	}
}

type Model struct {
	ID        string                `gorm:"primaryKey;type:varchar(20)" json:"id,omitempty"`
	CreatedAt uint64                `gorm:"autoCreateTime:milli;not null;index:idx_created_at;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt uint64                `gorm:"autoUpdateTime:milli;not null;index:idx_updated_at;comment:更新时间" json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;index:idx_deleted_at;comment:删除时间" json:"deleted_at,omitempty"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.ID = node.Generate().String()
	return nil
}
