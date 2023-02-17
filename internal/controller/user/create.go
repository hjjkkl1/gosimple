package user

import (
	"github.com/gin-gonic/gin"
	"github.com/valyevo/gosimple/internal/pkg/core"
)

func (ctrl *userController) Create(ctx *gin.Context) {
	err := ctrl.b.User().Create()
	core.WriteResponse(ctx, err, nil)
}
