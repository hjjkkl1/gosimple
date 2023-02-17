package biz

import (
	"github.com/valyevo/gosimple/internal/biz/user"
	"github.com/valyevo/gosimple/internal/store"
)

// IBiz 定义业务层实现的方法
type IBiz interface {
	User() user.UserBiz
}

// biz 是 IBiz 的具体实现
type biz struct {
	store store.IStore
}

var _ IBiz = (*biz)(nil)

// NewBiz 创建 IBiz 类型实例
func NewBiz(store store.IStore) *biz {
	return &biz{store}
}

// Users 返回实现 UserBiz 接口实例
func (b *biz) User() user.UserBiz {
	return user.NewUserBiz(b.store)
}
