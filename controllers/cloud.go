package controllers

import "github.com/astaxie/beego"

type CloudController struct {
	beego.Controller
}

// @Title CreateInstance
// @Description create instance
// @Param	body
// @Success 200 {int}
// @Failure 403
// @router /create [post]
func (cloud *CloudController) createInstance(){
	cloud.Data["json"] = map[string]string{}
	cloud.ServeJSON()
}

// @Title RunInstances
// @Description run instances
// @Param body
// @Success 200
// @Failure 403
// @router /run [post]
func (cloud *CloudController) runInstances(){
	dataMap := make(map[string]string)
	dataMap["instances"] = "a,b,c"
	cloud.Data["json"] = dataMap
	cloud.ServeJSON()
}