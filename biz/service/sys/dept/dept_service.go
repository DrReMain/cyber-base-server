package sys_dept_service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"github.com/DrReMain/cyber-base-server/biz/common/errc"
	"github.com/DrReMain/cyber-base-server/biz/common/pagi"
	"github.com/DrReMain/cyber-base-server/biz/dal/sys_model"
	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/sys/dept"
	cutils_default "github.com/DrReMain/cyber-base-server/cyber/utils/default"
)

type Service struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewService(ctx context.Context, c *app.RequestContext) *Service {
	return &Service{ctx, c}
}

func (s *Service) CreateDept(req *dept.CreateDeptReq) (err error) {
	_, err = sys_model.QueryByDeptName(*req.DeptName)
	if err == nil {
		err = errc.AlreadyExistErr
		return
	}

	m := &sys_model.SysDept{
		DeptName: req.DeptName,
		Remark:   req.Remark,
	}
	err = sys_model.CreateDept(m)
	return
}

func (s *Service) UpdateDept(req *dept.UpdateDeptReq) (err error) {
	_, err = sys_model.QueryDeptItem(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errc.NotExistErr
		return
	}

	m := &sys_model.SysDept{
		DeptName: req.DeptName,
		Remark:   req.Remark,
	}
	err = sys_model.UpdateDept(req.ID, m)
	return
}

func (s *Service) DeleteDept(req *dept.DeleteDeptReq) (err error) {
	_, err = sys_model.QueryDeptItem(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errc.NotExistErr
		return
	}

	err = sys_model.DeleteDept(req.ID)
	return
}

func (s *Service) QueryAllDept(req *dept.QueryAllDeptReq) (list *[]sys_model.SysDept, err error) {
	list, err = sys_model.QueryDeptAll(
		cutils_default.String(req.DeptName),
		req.CreatedAt,
	)
	return
}

func (s *Service) QueryListDept(req *dept.QueryListDeptReq) (list *[]sys_model.SysDept, p *pagi.Pagi, err error) {
	list, total, more, num, size, err := sys_model.QueryDeptList(
		cutils_default.Int(req.PageNum, 1),
		cutils_default.Int(req.PageSize, 10),
		cutils_default.String(req.DeptName),
	)
	p = pagi.NewPagi(total, more, num, size)
	return
}

func (s *Service) QueryItemDept(req *dept.QueryItemDeptReq) (item *sys_model.SysDept, err error) {
	item, err = sys_model.QueryDeptItem(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errc.NotExistErr
		return
	}
	return
}
