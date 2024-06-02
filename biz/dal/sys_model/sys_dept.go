package sys_model

import (
	"github.com/DrReMain/cyber-base-server/biz/dal"
	"github.com/DrReMain/cyber-base-server/cyber"
)

type SysDept struct {
	dal.Model
	DeptName *string `gorm:"type:varchar(100);not null;unique;comment:部门名称"`
	Remark   *string `gorm:"type:varchar(500);comment:备注"`
}

func (*SysDept) TableName() string {
	return "sys_dept"
}

func CreateDept(d *SysDept) (err error) {
	err = cyber.DB.Create(d).Error
	return
}

func UpdateDept(id string, d *SysDept) (err error) {
	err = cyber.DB.Model(d).Where("id = ?", id).Updates(d).Error
	return
}

func DeleteDept(id string) (err error) {
	err = cyber.DB.Where("id = ?", id).Delete(&SysDept{}).Error
	return
}

func QueryDeptAll(DeptName string) (list []SysDept, err error) {
	err = cyber.DB.Model(&SysDept{}).Where("dept_name LIKE ?", "%"+DeptName+"%").Find(&list).Error
	return
}

func QueryDeptList(pageNum, pageSize int, DeptName string) (list []SysDept, total int64, more bool, num, size int, err error) {
	num, size = pageNum, pageSize
	query := cyber.DB.Model(&SysDept{}).Where("dept_name LIKE ?", "%"+DeptName+"%")
	query.Count(&total)
	query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list)
	more = total >= int64(pageNum*pageSize)
	return
}

func QueryDeptItem(id string) (item SysDept, err error) {
	err = cyber.DB.Where("id = ?", id).First(&item).Error
	return
}
