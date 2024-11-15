package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	v1 "lite-frame/apis/v1"
	"lite-frame/config"
	"time"
)

var DB *gorm.DB

func InitMysql() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.MySql.Username, cfg.MySql.Password, cfg.MySql.Host, cfg.MySql.Port, cfg.MySql.Database)
	mysqlDb, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}
	hsqldb := mysqlDb.DB()
	hsqldb.SetMaxOpenConns(30)
	hsqldb.SetConnMaxLifetime(2 * time.Second)
	DB = mysqlDb

	return nil
}

func ApplyPageSql(db *gorm.DB, page v1.Page) (int64, error) {

	var count int64
	if page.PageSize > 0 {
		db = db.Limit(page.PageNumber).Offset((page.PageNumber - 1) * page.PageSize)
	} else {
		db = db.Offset(0)
	}
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	if page.OrderBy != "" {
		db = db.Order(page.PackOrderSql())
	}
	return count, nil
}
