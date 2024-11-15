package db

import (
	"encoding/json"
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

func (s *userStore) List(selector v1.UserSelector, page v1.Page) ([]*v1.User, int64, error) {
	var users []model.User
	db := s.db.Model(&model.User{})
	if selector.Name != "" {
		db.Where("name = ?", selector.Name)
	}
	count, err := dao.ApplyPageSql(db, page)
	if err != nil {
		return nil, 0, err
	}
	db.Find(&users)
	rows := make([]*v1.User, len(users))
	for i, user := range users {
		rows[i] = innerUserToV1(user)
	}
	return rows, count, nil
}

func innerUserToV1(user model.User) *v1.User {
	inByte, _ := json.Marshal(user)
	outer := &v1.User{}
	json.Unmarshal(inByte, outer)
	return outer
}

func v1UserToInner(user v1.User) *model.User {
	inByte, _ := json.Marshal(user)
	outer := &model.User{}
	json.Unmarshal(inByte, outer)
	return outer
}
