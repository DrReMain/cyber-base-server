package errc

import (
	"errors"
)

type ErrC struct {
	ErrCode string
	ErrMsg  string
}

func NewErrC(errCode string, errMsg string) ErrC {
	return ErrC{errCode, errMsg}
}

func (e ErrC) Error() string {
	return e.ErrMsg
}
func (e ErrC) Code() string {
	return e.ErrCode
}

func (e ErrC) WithMsg(errMsg string) ErrC {
	e.ErrMsg = errMsg
	return e
}

var (
	InternalErr            = NewErrC(InternalErrCode, "内部错误")
	ParamsInvalidErr       = NewErrC(ParamsInvalidErrCode, "参数错误")
	AuthorizationFailedErr = NewErrC(AuthorizationFailedErrCode, "权限不足")
	AlreadyExistErr        = NewErrC(AlreadyExistErrCode, "数据已存在")
	NotExistErr            = NewErrC(NotExistErrCode, "数据不存在")
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
