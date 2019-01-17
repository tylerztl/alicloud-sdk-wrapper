package controllers

import "github.com/astaxie/beego"

type InstanceController struct {
	beego.Controller
}

// @Title CreateInstance
// @Description create instance
// @Param	body
// @Success 200 {int}
// @Failure 403
// @router /create [post]
func (instance *InstanceController) createInstance(){

}

// @Title RunInstances
// @Description run instances
// @Param body
// @Success 200
// @Failure 403
// @router /run [post]
func (instance *InstanceController) runInstances(){

}