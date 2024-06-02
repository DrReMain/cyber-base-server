// Code generated by hertz generator.

package sys

import (
	"context"

	"github.com/DrReMain/cyber-base-server/biz/common/res"
	"github.com/DrReMain/cyber-base-server/biz/dal/sys_model"
	sys "github.com/DrReMain/cyber-base-server/biz/hertz_gen/sys"
	cutils_default "github.com/DrReMain/cyber-base-server/cyber/utils/default"

	"github.com/cloudwego/hertz/pkg/app"
)

// CreateDept .
// @router /v1/sys/dept/create [POST]
func CreateDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.CreateDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.CreateDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	m := &sys_model.SysDept{
		DeptName: req.DeptName,
		Remark:   req.Remark,
	}
	err = sys_model.CreateDept(m)
	if err != nil {
		res.InternalFail(c, &sys.CreateDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err, res.Json(m))
		return
	}

	res.Success(c, &sys.CreateDeptRes{
		Base:   res.BaseSuccess(),
		Result: nil,
	})
}

// UpdateDept .
// @router /v1/sys/dept/:id [PUT]
func UpdateDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.UpdateDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.UpdateDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	m := &sys_model.SysDept{
		DeptName: req.DeptName,
		Remark:   req.Remark,
	}
	err = sys_model.UpdateDept(req.ID, m)
	if err != nil {
		res.InternalFail(c, &sys.UpdateDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err, res.Json(m))
		return
	}

	res.Success(c, &sys.UpdateDeptRes{
		Base:   res.BaseSuccess(),
		Result: nil,
	})
}

// DeleteDept .
// @router /v1/sys/dept/:id [DELETE]
func DeleteDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.DeleteDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.DeleteDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	err = sys_model.DeleteDept(req.ID)
	if err != nil {
		res.InternalFail(c, &sys.DeleteDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err, req.ID)
		return
	}

	res.Success(c, &sys.DeleteDeptRes{
		Base:   res.BaseSuccess(),
		Result: nil,
	})
}

// QueryAllDept .
// @router /v1/sys/dept/list_all [GET]
func QueryAllDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryAllDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.QueryAllDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	l, err := sys_model.QueryDeptAll(cutils_default.String(req.DeptName))
	if err != nil {
		res.InternalFail(c, &sys.DeleteDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err)
		return
	}

	result := make([]*sys.Dept, 0, len(l))
	for _, dept := range l {
		result = append(result, &sys.Dept{
			ID:       dept.ID,
			DeptName: cutils_default.String(dept.DeptName),
			Remark:   cutils_default.String(dept.Remark),
		})
	}

	res.Success(c, &sys.QueryAllDeptRes{
		Base:   res.BaseSuccess(),
		Result: result,
	})
}

// QueryListDept .
// @router /v1/sys/dept/list [GET]
func QueryListDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryListDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.QueryListDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	l, t, m, n, s, err := sys_model.QueryDeptList(
		cutils_default.Int(req.PageNum, 1),
		cutils_default.Int(req.PageSize, 10),
		cutils_default.String(req.DeptName),
	)
	if err != nil {
		res.InternalFail(c, &sys.QueryListDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err)
		return
	}

	list := make([]*sys.Dept, 0, len(l))
	for _, dept := range l {
		list = append(list, &sys.Dept{
			ID:       dept.ID,
			DeptName: cutils_default.String(dept.DeptName),
			Remark:   cutils_default.String(dept.Remark),
		})
	}

	res.Success(c, &sys.QueryListDeptRes{
		Base: res.BaseSuccess(),
		Result: &sys.PResult{
			P:    res.P(t, m, n, s),
			List: list,
		},
	})
}

// QueryItemDept .
// @router /v1/sys/dept/:id [GET]
func QueryItemDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryItemDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		res.ValidateFail(c, &sys.QueryItemDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}

	item, err := sys_model.QueryDeptItem(req.ID)
	if err != nil {
		res.InternalFail(c, &sys.QueryItemDeptRes{
			Base:   res.BaseInternalFail(),
			Result: nil,
		}, err)
		return
	}

	res.Success(c, &sys.QueryItemDeptRes{
		Base: res.BaseSuccess(),
		Result: &sys.Dept{
			ID:       item.ID,
			DeptName: cutils_default.String(item.DeptName),
			Remark:   cutils_default.String(item.Remark),
		},
	})
}
