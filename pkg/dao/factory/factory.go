package factory

import (
	"lite-frame/pkg/dao"
	"lite-frame/pkg/dao/db"
)

func InitDB() {
	if err := dao.InitMysql(); err != nil {
		panic(err)
	}
	GetDao()
}

func GetDao() dao.Dao {
	return dao.Dao{
		User: db.InitUserStore(dao.DB),
	}
}
