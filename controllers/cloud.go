package controllers

import "github.com/astaxie/beego"

type CloudController struct {
	beego.Controller
}

// @Title CreateVPC
// @Description create VPC
// @Param	body	body 	commons.CreateVpcRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /vpc [post]
func (cloud *CloudController) CreateVPC(){
	cloud.Data["json"] = map[string]string{}
	cloud.ServeJSON()
}

// @Title CreateVSwitch
// @Description create VSwitch
// @Param	body	body 	commons.CreateVSwitchRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /vswitch [post]
func (cloud *CloudController) CreateVSwitch(){
	cloud.Data["json"] = map[string]string{}
	cloud.ServeJSON()
}


// @Title CreateSecurityGroup
// @Description create SecurityGroup
// @Param	body	body 	commons.CreateSecurityGroupRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /securitygroup [post]
func (cloud *CloudController) CreateSecurityGroup(){
	cloud.Data["json"] = map[string]string{}
	cloud.ServeJSON()
}



// @Title RunInstances
// @Description run instances
// @Param	body	body 	commons.RunInstancesRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /runinstances [post]
func (cloud *CloudController) RunInstances(){
	dataMap := make(map[string]string)
	dataMap["instances"] = "a,b,c"
	cloud.Data["json"] = dataMap
	cloud.ServeJSON()
}