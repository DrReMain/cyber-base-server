package errc

import (
	"errors"
	"fmt"

	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/code"
)

type ErrC struct {
	ErrCode code.Code
	ErrMsg  string
}

func NewErrC(errCode code.Code, errMsg string) ErrC {
	return ErrC{errCode, errMsg}
}

func (e ErrC) Error() string {
	return fmt.Sprintf("%s", e.ErrMsg)
}

func (e ErrC) WithMsg(errMsg string) ErrC {
	e.ErrMsg = errMsg
	return e
}

var (
	InternalErr            = NewErrC(code.Code_InternalErr, "内部错误")
	ParamsInvalidErr       = NewErrC(code.Code_ParamsInvalidErr, "参数错误")
	AuthorizationFailedErr = NewErrC(code.Code_AuthorizationFailedErr, "权限不足")
	AlreadyExistErr        = NewErrC(code.Code_AlreadyExistErr, "数据已存在")
	NotExistErr            = NewErrC(code.Code_NotExistErr, "数据不存在")
)

func ConvertInternal(err error) ErrC {
	e := Convert(err)
	e.ErrMsg = err.Error()
	return e
}

func Convert(err error) ErrC {
	e := ErrC{}
	if errors.As(err, &e) {
		return e
	}
	return InternalErr
}
