package dao

import v1 "lite-frame/apis/v1"

type UserInterface interface {
	Create(user v1.User) error
	List(selector v1.UserSelector, page v1.Page) ([]*v1.User, int64, error)
}

type Dao struct {
	User UserInterface
}
