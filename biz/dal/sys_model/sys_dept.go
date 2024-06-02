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

func QueryDeptAll(deptName string, CreatedAt []int64) (list *[]SysDept, err error) {
	list = &[]SysDept{}

	query := cyber.DB.Model(&SysDept{}).Where("dept_name LIKE ?", "%"+deptName+"%")
	if len(CreatedAt) == 2 {
		query = query.Where("created_at BETWEEN ? AND ?", CreatedAt[0], CreatedAt[1])
	}
	err = query.Find(list).Error
	return
}

func QueryDeptList(pageNum, pageSize int, deptName string) (list *[]SysDept, total int64, more bool, num, size int, err error) {
	list = &[]SysDept{}

	query := cyber.DB.Model(&SysDept{}).Where("dept_name LIKE ?", "%"+deptName+"%")
	err = query.Count(&total).Error
	err = query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(list).Error
	more = total >= int64(pageNum*pageSize)
	num, size = pageNum, pageSize
	return
}

func QueryDeptItem(id string) (item *SysDept, err error) {
	item = &SysDept{}
	err = cyber.DB.Where("id = ?", id).First(item).Error
	return
}

func QueryByDeptName(deptName string) (item *SysDept, err error) {
	item = &SysDept{}
	err = cyber.DB.Where("dept_name = ?", deptName).First(item).Error
	return
}
