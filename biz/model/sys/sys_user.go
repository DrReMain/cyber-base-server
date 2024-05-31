package sys

import (
	"github.com/gofrs/uuid/v5"

	"github.com/DrReMain/cyber-base-server/biz/model"
)

type SysUser struct {
	model.Model
	UUID     uuid.UUID `gorm:"index;"`
	Mobile   string    `gorm:"type:varchar(20);not null;unique;index:idx_mobile_deleted_at;comment:登录手机号"`
	Password string    `gorm:"type:varchar(100);not null;comment:密码"`
}

func (u *SysUser) TableName() string {
	return "sys_user"
}
