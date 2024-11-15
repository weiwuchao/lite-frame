package service

import (
	v1 "lite-frame/apis/v1"
	"lite-frame/pkg/dao/factory"
)

func CreateUser(user v1.User) error {
	return factory.GetDao().User.Create(user)
}
