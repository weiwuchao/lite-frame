package service

import (
	v1 "lite-frame/apis/v1"
	"lite-frame/pkg/dao/factory"
)

func CreateUser(user v1.User) error {
	return factory.GetDao().User.Create(user)
}

func ListUser(selector v1.UserSelector, page v1.Page) ([]*v1.User, int64, error) {
	return factory.GetDao().User.List(selector, page)
}
