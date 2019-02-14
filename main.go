package main

import (
	"encoding/json"
	_ "zig-cloud/database"
	_ "zig-cloud/routers"
	_ "zig-cloud/services/cloudprovider/alicloud"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	initLogger()
}

func initLogger() {
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("logPath")
	logLevel, err := beego.AppConfig.Int("logLevel")
	if nil != err {
		logLevel = beego.LevelDebug
	}
	config["level"] = logLevel

	configStr, err := json.Marshal(config)
	if nil != err {
		panic(err)
	}
	err = beego.SetLogger(logs.AdapterFile, string(configStr))
	beego.SetLogFuncCall(true)
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
