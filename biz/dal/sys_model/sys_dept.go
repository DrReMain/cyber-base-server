package sys_model

import (
	"github.com/DrReMain/cyber-base-server/biz/dal"
	"github.com/DrReMain/cyber-base-server/cyber"
	"github.com/gofrs/uuid/v5"
)

type SysDept struct {
	dal.Model
	UUID     uuid.UUID `gorm:"index;"`
	DeptName string    `gorm:"type:varchar(100);not null;unique;index:idx_deptname_deleted_at;comment:部门名称"`
	Remark   string    `gorm:"type:varchar(500);comment:备注"`
}

func (*SysDept) TableName() string {
	return "sys_dept"
}

func CreateDept(d *SysDept) error {
	return cyber.DB.Create(&d).Error
}

func UpdateDept(d *SysDept, id uint64) (err error) {
	err = cyber.DB.Model(&d).Where("id = ?", id).Updates(d).Error
	return
}

func DeleteDept(id uint64) (err error) {
	err = cyber.DB.Where("id = ?", id).Delete(&SysDept{}).Error
	return
}

func QueryDept(id uint64) (item *SysDept, err error) {
	err = cyber.DB.Where("id = ?", id).First(item).Error
	return
}

func QueryDeptList(pageNum, pageSize int, deptName string) (list []SysDept, num int, size int, total int64, more bool, err error) {
	num, size = pageNum, pageSize
	query := cyber.DB.Model(&SysDept{}).Where("dept_name LIKE ?", "%"+deptName+"%")
	query.Count(&total)
	query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list)
	more = total >= int64(pageNum*pageSize)
	return
}

func QueryDeptListAll(deptName string) (list []SysDept, err error) {
	err = cyber.DB.Model(&SysDept{}).Where("dept_name LIKE = ?", "%"+deptName+"%").Find(&list).Error
	return
}
