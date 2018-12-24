package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"zig-cloud/models"
)

func RegisterSQLite() {
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "sqlite3", "zigcloud.db", maxIdle, maxConn)
	RegisterModels()
	orm.RunSyncdb("default", false, true)
}

func RegisterModels() {
	orm.RegisterModel(new(models.User))
}
