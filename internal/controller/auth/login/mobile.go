package login

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/valyevo/gosimple/internal/pkg/core"
	"github.com/valyevo/gosimple/internal/pkg/errcode"
	"github.com/valyevo/gosimple/internal/request/login"
)

func (ctrl *loginController) Mobile(ctx *gin.Context) {
	var r login.MobileLoginRequest
	// 绑定参数
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errcode.ErrBind, nil)
		return
	}
	// 验证参数
	_ = r
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(ctx, errcode.ErrInvalidParameter.SetMsg(err.Error()), nil)
		return
	}

	// TODO: 校验数据库

	return
}
