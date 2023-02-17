package user

import (
	"github.com/valyevo/gosimple/internal/pkg/errcode"
	"github.com/valyevo/gosimple/internal/store"
)

// UserBiz 定义 User 业务模块实现的方法
type UserBiz interface {
	Create() error // 创建用户
	Index()        // 用户列表
}

// userBiz UserBiz 的具体实现
type userBiz struct {
	store store.IStore
}

// NewUserBiz 实例化 UserBiz 接口实例
func NewUserBiz(store store.IStore) *userBiz {
	return &userBiz{store}
}

// Create 是 UserBiz 接口 `Create` 方法具体实现
func (b *userBiz) Create() error {
	err := b.store.User().Create()

	if err != nil {
		return errcode.ServerError
	}

	return nil
}

// Index 是 UserBiz 接口 `Index` 方法具体实现
func (b *userBiz) Index() {

}
