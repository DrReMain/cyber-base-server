package sys_model

import (
	"github.com/DrReMain/cyber-base-server/biz/dal"
)

type SysUser struct {
	dal.Model
	Mobile   string `gorm:"type:varchar(11);not null;unique;comment:登录手机号"`
	Password string `gorm:"type:varchar(100);not null;comment:密码"`
}

func (u *SysUser) TableName() string {
	return "sys_user"
}
