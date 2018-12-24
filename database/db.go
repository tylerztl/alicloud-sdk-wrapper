package database

import "github.com/astaxie/beego/orm"

func GetOrmer() orm.Ormer {
	return orm.NewOrm()
}
