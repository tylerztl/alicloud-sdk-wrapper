package database

import (
	"zig-cloud/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func RegisterSQLite() {
	maxIdle := 30
	maxConn := 30
	err := orm.RegisterDataBase("default", "sqlite3", "zigcloud.db", maxIdle, maxConn)
	if nil != err {
		panic(err)
	}
	RegisterModels()
	err = orm.RunSyncdb("default", false, true)
	if nil != err {
		panic(err)
	}
}

func RegisterModels() {
	orm.RegisterModel(new(models.User))
}
