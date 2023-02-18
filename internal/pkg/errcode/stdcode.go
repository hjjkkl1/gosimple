package errcode

const defref = "https://github.com/valyevo/gosimple/tree/main/docs"
const nullref = "无"

var (
	OK                  = new(200, "0", "ok", nullref)
	ServerError         = new(500, "100001", "服务器内部错误", nullref)
	ErrNotFound         = new(404, "100002", "找不到资源", nullref)
	ErrBind             = new(400, "100003", "缺省参数", defref)
	ErrInvalidParameter = new(400, "100003", "无效的参数", defref)
)
