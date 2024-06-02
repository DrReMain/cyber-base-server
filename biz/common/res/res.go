package res

import (
	"context"
	"errors"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/base"
	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/code"
)

type Res struct {
	ctx context.Context
	c   *app.RequestContext
	req any
}

func NewRes(ctx context.Context, c *app.RequestContext, req any) *Res {
	return &Res{ctx, c, req}
}

func (r *Res) ValidateFail(err error, o any) {
	hlog.Infof("[%s]: %s \r\n%s", code.Code_ParamsInvalid.String(), err, Json(r.req))
	r.c.JSON(consts.StatusOK, o)
}

func (r *Res) InternalFail(err error, o any, a ...any) {
	hlog.Infof("[%s]: %s \r\n%s", code.Code_DBError.String(), err, Json(a))
	r.c.JSON(consts.StatusInternalServerError, o)
}

func (r *Res) Success(o any) {
	r.c.JSON(consts.StatusOK, o)
}

func Base(success bool, code code.Code, rest ...any) *base.Base {
	var msg string
	if len(rest) > 0 {
		if err, ok := rest[0].(error); ok {
			msg = err.Error()
		}
	} else {
		msg = "OK"
	}
	return &base.Base{
		T:       time.Now().UnixMilli(),
		Success: success,
		Code:    code,
		Msg:     msg,
	}
}
func BaseSuccess() *base.Base {
	return Base(true, code.Code_Success)
}
func BaseValidateFail(err error) *base.Base {
	return Base(false, code.Code_ParamsInvalid, err)
}
func BaseInternalFail() *base.Base {
	return Base(false, code.Code_DBError, errors.New("服务器异常"))
}

func Json(o any) string {
	j, _ := sonic.MarshalString(o)
	return j
}
