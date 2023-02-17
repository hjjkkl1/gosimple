package store

import (
	"errors"

	"gorm.io/gorm"
)

type UserStore interface {
	Create() error
	Get() error
}

type user struct {
	db *gorm.DB
}

func newUser(db *gorm.DB) *user {
	return &user{db}
}

var _ UserStore = (*user)(nil)

func (u *user) Create() error {
	return errors.New("create failed")
}
func (u *user) Get() error {
	return nil
}
