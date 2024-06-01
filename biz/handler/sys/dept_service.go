// Code generated by hertz generator.

package sys

import (
	"context"
	"github.com/DrReMain/cyber-base-server/biz/common/res"
	"github.com/DrReMain/cyber-base-server/biz/dal/sys_model"
	sys "github.com/DrReMain/cyber-base-server/biz/hertz_gen/sys"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gofrs/uuid/v5"
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
	{
		m := &sys_model.SysDept{
			UUID:     uuid.Must(uuid.NewV4()),
			DeptName: *req.DeptName,
			Remark:   *req.Remark,
		}
		err = sys_model.CreateDept(m)
		if err != nil {
			o := &sys.CreateDeptRes{
				Base:   res.BaseInternalFail(),
				Result: nil,
			}
			res.InternalFail(c, o, err, res.Json(m))
			return
		}
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
		res.ValidateFail(c, &sys.CreateDeptRes{
			Base:   res.BaseValidateFail(err),
			Result: nil,
		}, err, res.Json(req))
		return
	}
	{
		m := &sys_model.SysDept{
			DeptName: *req.DeptName,
			Remark:   *req.Remark,
		}
		err = sys_model.UpdateDept(m, uint64(req.ID))
		if err != nil {
			o := &sys.CreateDeptRes{
				Base:   res.BaseInternalFail(),
				Result: nil,
			}
			res.InternalFail(c, o, err, res.Json(m))
			return
		}
	}
	res.Success(c, &sys.CreateDeptRes{
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
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys.DeleteDeptRes)

	c.JSON(consts.StatusOK, resp)
}

// QueryAllDept .
// @router /v1/sys/dept/list_all [GET]
func QueryAllDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryAllDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys.QueryAllDeptRes)

	c.JSON(consts.StatusOK, resp)
}

// QueryListDept .
// @router /v1/sys/dept/list [GET]
func QueryListDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryListDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys.QueryListDeptRes)

	c.JSON(consts.StatusOK, resp)
}

// QueryItemDept .
// @router /v1/sys/dept/:id [GET]
func QueryItemDept(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.QueryItemDeptReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sys.QueryItemDeptRes)

	c.JSON(consts.StatusOK, resp)
}
