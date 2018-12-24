package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/mattn/go-sqlite3"
	"zig-cloud/database"
	_ "zig-cloud/routers"
)

func init() {
	database.RegisterSQLite()
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:   []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
