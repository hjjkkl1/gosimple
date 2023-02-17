package user

import (
	"github.com/valyevo/gosimple/internal/biz"
	"github.com/valyevo/gosimple/internal/store"
)

type userController struct {
	b biz.IBiz
}

func New(store store.IStore) *userController {
	return &userController{b: biz.NewBiz(store)}
}
