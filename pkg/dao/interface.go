package dao

import v1 "lite-frame/apis/v1"

type UserInterface interface {
	Create(user v1.User) error
}

type Dao struct {
	User UserInterface
}