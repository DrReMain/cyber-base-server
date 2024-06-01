package res

import (
	"errors"
	"github.com/bytedance/sonic"
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
func BaseSuccess() *base.Base {
	return Base(true, code.Code_Success)
}
func BaseValidateFail(err error) *base.Base {
	return Base(false, code.Code_ParamsInvalid, err)
}
func BaseInternalFail() *base.Base {
	return Base(false, code.Code_DBError, errors.New("服务器错误"))
}

func Success(c *app.RequestContext, o any) {
	c.JSON(consts.StatusOK, o)
}
func ValidateFail(c *app.RequestContext, o any, err error, p ...string) {
	hlog.Infof("[%s]: %s %s\n", code.Code_ParamsInvalid.String(), err, p)
	c.JSON(consts.StatusOK, o)
}
func InternalFail(c *app.RequestContext, o any, err error, p ...string) {
	hlog.Infof("[%s]: %s %s\n", code.Code_DBError.String(), err, p)
	c.JSON(consts.StatusInternalServerError, o)
}

func Json(o any) string {
	j, _ := sonic.Marshal(o)
	return string(j)
}
