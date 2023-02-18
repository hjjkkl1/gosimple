package errcode

import "fmt"

// Errcode 定制错误包
type Errcode struct {
	Http int    // HTTP Code
	Code string // Code
	Msg  string // Message
	Ref  string // Ref
}

// New 实例化一个错误包
func new(httpStatus int, appcode, msg, ref string) *Errcode {
	return &Errcode{httpStatus, appcode, msg, ref}
}

// Error 实现 error 标准包接口
func (e *Errcode) Error() string {
	return e.Msg
}

// SetMsg 设置错误信息
func (e *Errcode) SetMsg(fomatString string, args ...interface{}) *Errcode {
	e.Msg = fmt.Sprintf(fomatString, args...)
	return e
}

// DecodeE 解析错误
func DecodeE(err error) (int, string, string, string) {
	if err == nil {
		return OK.Http, OK.Code, OK.Msg, OK.Ref
	}

	// 判断错误类型
	switch typ := err.(type) {
	case *Errcode: // 类型是 *Errcode
		return typ.Http, typ.Code, typ.Msg, typ.Ref
	}

	return ServerError.Http, ServerError.Code, ServerError.Msg, ServerError.Ref
}
