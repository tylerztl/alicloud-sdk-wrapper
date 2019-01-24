package controllers

import (
	"encoding/json"
	"zig-cloud/commons"
	"zig-cloud/services"

	"github.com/astaxie/beego"
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
func (cloud *CloudController) CreateVPC() {
	createVpcRequest := new(commons.CreateVpcRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createVpcRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createVpcResponse, err := cloudProvider.CreateVpc(createVpcRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["VpcId"] = createVpcResponse.VpcId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title CreateVSwitch
// @Description create VSwitch
// @Param	body	body 	commons.CreateVSwitchRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /vswitch [post]
func (cloud *CloudController) CreateVSwitch() {
	createVSwitchRequest := new(commons.CreateVSwitchRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createVSwitchRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createVSwitchResponse, err := cloudProvider.CreateVSwitch(createVSwitchRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["VSwitchId"] = createVSwitchResponse.VSwitchId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title CreateSecurityGroup
// @Description create SecurityGroup
// @Param	body	body 	commons.CreateSecurityGroupRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /securitygroup [post]
func (cloud *CloudController) CreateSecurityGroup() {
	createSecurityGroupRequest := new(commons.CreateSecurityGroupRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createSecurityGroupRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	createSecurityGroupResponse, err := cloudProvider.CreateSecurityGroup(createSecurityGroupRequest)
	if err == nil {
		responseData := make(map[string]string)
		responseData["SecurityGroupId"] = createSecurityGroupResponse.SecurityGroupId
		cloud.Data["json"] = responseData
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title RunInstances
// @Description run instances
// @Param	body	body 	commons.RunInstancesRequest	true 	"body content"
// @Success 200 {}
// @Failure 403
// @router /runinstances [post]
func (cloud *CloudController) RunInstances() {
	runInstancesRequest := new(commons.RunInstancesRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, runInstancesRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	runInstancesResponse, err := cloudProvider.RunInstances(runInstancesRequest)
	if err == nil {
		cloud.Data["json"] = runInstancesResponse.Instances
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title DescribeRegions
// @Description Query available regions
// @Success 200 {}
// @Failure 403
// @router /regions [get]
func (cloud *CloudController) DescribeRegions() {
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	describeRegionsResponse, err := cloudProvider.DescribeRegions()
	if err == nil {
		cloud.Data["json"] = describeRegionsResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title DescribeZones
// @Description Query the available zones for the specified region
// @Param	regionId		query 	string	true		"The regionId for zones"
// @Success 200 {}
// @Failure 403
// @router /zones [get]
func (cloud *CloudController) DescribeZones() {
	regionId := cloud.Ctx.Input.Query("regionId")
	cloudProvider := services.GetProviderByType(commons.CloudProviderAliCloud)
	describeZonesResponse, err := cloudProvider.DescribeZones(regionId)
	if err == nil {
		cloud.Data["json"] = describeZonesResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}
