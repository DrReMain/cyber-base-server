package res

import (
	"time"

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
