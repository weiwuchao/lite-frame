package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:64"`
	Password string `gorm:"size:64"`
	Email    string `gorm:"size:64"`
}
