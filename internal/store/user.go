package store

import "gorm.io/gorm"

type UserStore interface {
	Create() error
	Get() error
}

type user struct {
	*gorm.DB
}

func newUser(db *gorm.DB) *user {
	return &user{db}
}

var _ UserStore = (*user)(nil)

func (u *user) Create() error {
	return nil
}
func (u *user) Get() error {
	return nil
}
