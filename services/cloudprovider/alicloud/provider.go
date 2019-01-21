package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/helpers"
	"zig-cloud/commons"
	"zig-cloud/services"
	"github.com/astaxie/beego"
	"errors"
)

type CloudProvider struct {
	Config *services.AliCloudConfig
}

func (provider *CloudProvider) ConfigureClient(config *services.AliCloudConfig) {
	provider.Config = config
}

func (provider *CloudProvider) CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	vpcRequest := ecs.CreateCreateVpcRequest()
	if request.VpcName == commons.ValueEmptry {
		vpcRequest.VpcName = commons.AliCloudVPCName
	}else {
		vpcRequest.VpcName = request.VpcName
	}
	if request.RegionId == commons.ValueEmptry {
		vpcRequest.RegionId = commons.AliCloudRegionId
	}else {
		vpcRequest.RegionId = request.RegionId
	}
	if request.CidrBlock == commons.ValueEmptry {
		vpcRequest.CidrBlock = commons.AliCloudVPCCidrBlock
	}else {
		vpcRequest.CidrBlock = request.CidrBlock
	}
	vpcResponse, err := client.CreateVpc(vpcRequest)
	if err == nil {
		return &commons.CreateVpcResponse{VpcId:vpcResponse.VpcId},nil
	}else {
		return nil,err
	}
}

func (provider *CloudProvider) CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	switchRequest := ecs.CreateCreateVSwitchRequest()
	if request.VpcId == commons.ValueEmptry {
		return nil, errors.New(commons.AliCloudErrorMessageVPC)
	}else {
		switchRequest.VpcId = request.VpcId
	}
	if request.VSwitchName == commons.ValueEmptry {
		switchRequest.VSwitchName = commons.AliCloudVSwitchName
	}else {
		switchRequest.VSwitchName = request.VSwitchName
	}
	if request.ZoneId == commons.ValueEmptry {
		switchRequest.ZoneId = commons.AliCloudZoneId
	}else {
		switchRequest.ZoneId = request.ZoneId
	}
	if request.CidrBlock == commons.ValueEmptry {
		switchRequest.CidrBlock = commons.AliCloudSwitchCidrBlock
	}else {
		switchRequest.CidrBlock = request.CidrBlock
	}
	switchResponse, err := client.CreateVSwitch(switchRequest)
	if err == nil {
		return &commons.CreateVSwitchResponse{VSwitchId:switchResponse.VSwitchId}, nil
	}else {
		return nil, err
	}
}

func (provider *CloudProvider) CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error) {

	return nil,nil
}

func (provider *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	runInstancesRequest := ecs.CreateRunInstancesRequest()
	helpers.TransferValuesBetweenRequest(runInstancesRequest,request)
	response, err := client.RunInstances(runInstancesRequest)
	if err == nil {
		return helpers.GetRunInstancesResponse(response), nil
	} else {
		return nil, err
	}
}

func (provider *CloudProvider) GetClient() (*ecs.Client, error) {
	config := provider.Config
	client, err := ecs.NewClientWithAccessKey(config.RegionId, config.AccessKeyId, config.AccessKeySecret)
	if err == nil {
		return client, nil
	}else {
		beego.Error(err)
		return nil, err
	}
}