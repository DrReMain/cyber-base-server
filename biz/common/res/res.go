package res

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/base"
	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/code"
)

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

type parse interface {
	String() string
}

func ValidateFail(c *app.RequestContext, o any, code, req parse, err error) {
	hlog.Infof("[%s]: %s --> %s \n", code.String(), req.String(), err)
	c.JSON(consts.StatusOK, o)
}

func Success(c *app.RequestContext, o any) {
	c.JSON(consts.StatusOK, o)
}
