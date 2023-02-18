package login

import (
	"github.com/valyevo/gosimple/internal/biz"
	"github.com/valyevo/gosimple/internal/store"
)

// loginController 登录控制器
type loginController struct {
	b biz.IBiz
}

// NewLoginController 登录控制器实例
func NewLoginController(store store.IStore) *loginController {
	return &loginController{b: biz.NewBiz(store)}
}
