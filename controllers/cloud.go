package controllers

import "github.com/astaxie/beego"

type CloudController struct {
	beego.Controller
}

// @Title CreateInstance
// @Description create instance
// @Param	body	body 	commons.CreateInstanceRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /create [post]
func (cloud *CloudController) CreateInstance(){
	cloud.Data["json"] = map[string]string{}
	cloud.ServeJSON()
}

// @Title RunInstances
// @Description run instances
// @Param	body	body 	commons.RunInstancesRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /run [post]
func (cloud *CloudController) RunInstances(){
	dataMap := make(map[string]string)
	dataMap["instances"] = "a,b,c"
	cloud.Data["json"] = dataMap
	cloud.ServeJSON()
}