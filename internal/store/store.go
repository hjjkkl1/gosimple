package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	D    *datastore
)

// IStore 定义需要实现的方法
type IStore interface {
	User() UserStore
}

// datastore 是 IStore 的具体实现
type datastore struct {
	gdb *gorm.DB
}

// NewStore 实现 IStore
func NewDatastore(db *gorm.DB) *datastore {
	once.Do(func() {
		D = &datastore{db}
	})

	return D
}

var _ IStore = (*datastore)(nil)

// User 用户接口方法
func (ds *datastore) User() UserStore {
	return newUser(ds.gdb)
}
