package db

import (
	"github.com/jinzhu/gorm"
	v1 "lite-frame/apis/v1"
	"lite-frame/pkg/dao"
	"lite-frame/pkg/model"
	"log"
)

type userStore struct {
	db *gorm.DB
}

func InitUserStore(db *gorm.DB) dao.UserInterface {
	if err := db.AutoMigrate(&model.User{}).Error; err != nil {
		log.Fatalf("auto migrate user err: %v", err)
	}
	return &userStore{db: db}
}

func (s *userStore) Create(user v1.User) error {
	return s.db.Create(&user).Error
}
