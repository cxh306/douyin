package Error

import (
	"fmt"
)

const (
	SuccessCode             = 0
	EnPwdErrCode            = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

var (
	Success = NewErrNo(SuccessCode, "Success")

	EnPwdErr            = NewErrNo(EnPwdErrCode, "密码加密错误")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
)
