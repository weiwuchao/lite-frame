package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
