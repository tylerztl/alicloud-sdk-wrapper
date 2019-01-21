package controllers

import (
	"github.com/astaxie/beego"
	"zig-cloud/commons"
	"encoding/json"
	"zig-cloud/services"
)

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
	createVpcRequest := new(commons.CreateVpcRequest)
	json.Unmarshal(cloud.Ctx.Input.RequestBody, createVpcRequest)
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createVpcResponse, err := cloudProvider.CreateVpc(createVpcRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["VpcId"] = createVpcResponse.VpcId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403,err.Error())
	}
}

// @Title CreateVSwitch
// @Description create VSwitch
// @Param	body	body 	commons.CreateVSwitchRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /vswitch [post]
func (cloud *CloudController) CreateVSwitch(){
	createVSwitchRequest := new(commons.CreateVSwitchRequest)
	json.Unmarshal(cloud.Ctx.Input.RequestBody, createVSwitchRequest)
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createVSwitchResponse, err := cloudProvider.CreateVSwitch(createVSwitchRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["VSwitchId"] = createVSwitchResponse.VSwitchId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403,err.Error())
	}
}


// @Title CreateSecurityGroup
// @Description create SecurityGroup
// @Param	body	body 	commons.CreateSecurityGroupRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /securitygroup [post]
func (cloud *CloudController) CreateSecurityGroup(){
	createSecurityGroupRequest := new(commons.CreateSecurityGroupRequest)
	json.Unmarshal(cloud.Ctx.Input.RequestBody, createSecurityGroupRequest)
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createSecurityGroupResponse, err := cloudProvider.CreateSecurityGroup(createSecurityGroupRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["SecurityGroupId"] = createSecurityGroupResponse.SecurityGroupId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403,err.Error())
	}
}

// @Title RunInstances
// @Description run instances
// @Param	body	body 	commons.RunInstancesRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /runinstances [post]
func (cloud *CloudController) RunInstances(){
	runInstancesRequest := new(commons.RunInstancesRequest)
	json.Unmarshal(cloud.Ctx.Input.RequestBody, runInstancesRequest)
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	runInstancesResponse, err := cloudProvider.RunInstances(runInstancesRequest)
	if err == nil {
		cloud.Data["json"] = runInstancesResponse.Instances
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403,err.Error())
	}
}