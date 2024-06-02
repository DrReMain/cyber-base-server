package sys_model

import (
	"github.com/DrReMain/cyber-base-server/biz/dal"
	"github.com/DrReMain/cyber-base-server/cyber"
)

type SysUser struct {
	dal.Model
	Email    *string `gorm:"type:varchar(100);not null;uniqueIndex;comment:登录邮箱"`
	Password *string `gorm:"type:varchar(100);not null;comment:密码"`
	Mobile   *string `gorm:"type:varchar(11);uniqueIndex;comment:手机号"`
	Name     *string `gorm:"type:varchar(100);default:普通用户;comment:用户名称"`
	Avatar   *string `gorm:"type:varchar(200);default:https://raw.githubusercontent.com/DrReMain/DrReMain/main/README/go-original.svg;comment:用户头像"`
	Ban      int32   `gorm:"default:0;comment:用户禁用,0否,1是"`
}

func (*SysUser) TableName() string {
	return "sys_user"
}

func CreateUser(u *SysUser) (err error) {
	err = cyber.DB.Create(u).Error
	return
}

func UpdateUser(id string, u *SysUser) (err error) {
	err = cyber.DB.Model(u).Where("id = ?", id).Updates(u).Error
	return
}

func QueryByEmail(email string) (item *SysUser, err error) {
	item = &SysUser{}
	err = cyber.DB.Where("email = ?", email).First(&item).Error
	return
}
