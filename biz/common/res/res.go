package res

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/DrReMain/cyber-base-server/biz/common/errc"
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

func (r *Res) Fail(err error, o any) {
	if err == nil {
		return
	}
	e := errc.ConvertInternal(err)
	hlog.Infof("[%s]: %s\r\n%s", e.ErrCode.String(), e.Error(), Json(r.req))
	r.c.JSON(consts.StatusOK, o)
}
func (r *Res) Success(o any) {
	r.c.JSON(consts.StatusOK, o)
}

func Base(errRest ...error) *base.Base {
	var success = true
	var c = code.Code_Success
	var msg = "ok"

	if len(errRest) > 0 {
		if err, ok := errRest[0].(error); ok {
			e := errc.Convert(err)
			success = false
			c = e.ErrCode
			msg = e.Error()
		}
	}

	return &base.Base{
		T:       time.Now().UnixMilli(),
		Success: success,
		Code:    c,
		Msg:     msg,
	}
}

func Json(o any) string {
	j, _ := sonic.MarshalString(o)
	return j
}
