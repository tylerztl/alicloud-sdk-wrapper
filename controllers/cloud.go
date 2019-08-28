package controllers

import (
	"encoding/json"
	"alicloud-sdk-wrapper/commons"
	"alicloud-sdk-wrapper/services"

	"github.com/astaxie/beego"
)

type CloudController struct {
	Provider commons.CloudProvider
	beego.Controller
}

// @Title CreateVPC
// @Description create VPC
// @Param	body	body 	commons.CreateVpcRequest	true 	"body content"
// @Success 200 {object} commons.CreateVpcResponse
// @Failure 403
// @router /vpc [post]
func (cloud *CloudController) CreateVPC() {
	createVpcRequest := new(commons.CreateVpcRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createVpcRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(cloud.Provider)
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
// @Success 200 {object} commons.CreateVSwitchResponse
// @Failure 403
// @router /vswitch [post]
func (cloud *CloudController) CreateVSwitch() {
	createVSwitchRequest := new(commons.CreateVSwitchRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createVSwitchRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(cloud.Provider)
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
// @Success 200 {object} commons.CreateSecurityGroupResponse
// @Failure 403
// @router /securitygroup [post]
func (cloud *CloudController) CreateSecurityGroup() {
	createSecurityGroupRequest := new(commons.CreateSecurityGroupRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, createSecurityGroupRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(cloud.Provider)
	createSecurityGroupResponse, err := cloudProvider.CreateSecurityGroup(createSecurityGroupRequest)
	if err == nil {
		cloud.Data["json"] = createSecurityGroupResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title AuthorizeSecurityGroup
// @Description Add a security group entry direction rule
// @Param	body	body 	commons.AuthorizeSecurityGroupRequest	true 	"body content"
// @Success 200 {object} commons.AuthorizeSecurityGroupResponse
// @Failure 403
// @router /securitygroup/ingress-rule [post]
func (cloud *CloudController) AuthorizeSecurityGroup() {
	authorizeSecurityGroupRequest := new(commons.AuthorizeSecurityGroupRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, authorizeSecurityGroupRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(cloud.Provider)
	authorizeSecurityGroupResponse, err := cloudProvider.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
	if err == nil {
		cloud.Data["json"] = authorizeSecurityGroupResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title RunInstances
// @Description run instances
// @Param	body	body 	commons.RunInstancesRequest	true 	"body content"
// @Success 200 {object} commons.RunInstancesResponse
// @Failure 403
// @router /instances [post]
func (cloud *CloudController) RunInstances() {
	runInstancesRequest := new(commons.RunInstancesRequest)
	err := json.Unmarshal(cloud.Ctx.Input.RequestBody, runInstancesRequest)
	if nil != err {
		cloud.CustomAbort(403, err.Error())
	}
	cloudProvider := services.GetProviderByType(cloud.Provider)
	runInstancesResponse, err := cloudProvider.RunInstances(runInstancesRequest)
	if err == nil {
		cloud.Data["json"] = runInstancesResponse.Instances
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title DescribeInstances
// @Description Query for details of one or more instances
// @Param	regionId		query 	string	true		"The regionId for instances"
// @Param	instanceIds		query 	string	true		"The instanceIds for query"
// @Success 200 {object} commons.DescribeInstancesResponse
// @Failure 403
// @router /instances [get]
func (cloud *CloudController) DescribeInstances() {
	regionId := cloud.GetString("regionId")
	instanceIds := cloud.GetString("instanceIds")
	cloudProvider := services.GetProviderByType(cloud.Provider)
	describeRegionsResponse, err := cloudProvider.DescribeInstances(&commons.DescribeInstancesRequest{
		RegionId:    regionId,
		InstanceIds: instanceIds,
	})
	if err == nil {
		cloud.Data["json"] = describeRegionsResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title StopInstance
// @Description Stop an running instance
// @Param	instanceId		path 	string	true		"The instanceId for stop"
// @Success 200 {object} commons.StopInstanceResponse
// @Failure 403
// @router /instance/stop/:instanceId [get]
func (cloud *CloudController) StopInstance() {
	instanceId := cloud.GetString(":instanceId")
	cloudProvider := services.GetProviderByType(cloud.Provider)
	stopInstanceResponse, err := cloudProvider.StopInstance(instanceId)
	if err == nil {
		cloud.Data["json"] = stopInstanceResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}

// @Title DeleteInstance
// @Description Release a PostPaid instance or an expiring PrePaid instance
// @Param	instanceId		path 	string	true		"The instanceId for delete"
// @Success 200 {object} commons.DeleteInstanceResponse
// @Failure 403
// @router /instance/delete/:instanceId [get]
func (cloud *CloudController) DeleteInstance() {
	instanceId := cloud.GetString(":instanceId")
	cloudProvider := services.GetProviderByType(cloud.Provider)
	deleteInstanceResponse, err := cloudProvider.DeleteInstance(instanceId)
	if err == nil {
		cloud.Data["json"] = deleteInstanceResponse
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
	cloudProvider := services.GetProviderByType(cloud.Provider)
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
// @Success 200 {object} []commons.DescribeZonesResponse
// @Failure 403
// @router /zones [get]
func (cloud *CloudController) DescribeZones() {
	regionId := cloud.GetString("regionId")
	cloudProvider := services.GetProviderByType(cloud.Provider)
	describeZonesResponse, err := cloudProvider.DescribeZones(regionId)
	if err == nil {
		cloud.Data["json"] = describeZonesResponse
		cloud.ServeJSON()
	} else {
		cloud.CustomAbort(403, err.Error())
	}
}
