package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/helpers"
	"zig-cloud/commons"
	"zig-cloud/services"
	"github.com/astaxie/beego"
	"errors"
	"time"
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
	if request.VpcName == commons.ValueEmpty {
		vpcRequest.VpcName = commons.AliCloudVPCName
	}else {
		vpcRequest.VpcName = request.VpcName
	}
	if request.RegionId == commons.ValueEmpty {
		vpcRequest.RegionId = commons.AliCloudRegionId
	}else {
		vpcRequest.RegionId = request.RegionId
	}
	if request.CidrBlock == commons.ValueEmpty {
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
	if request.VpcId == commons.ValueEmpty {
		return nil, errors.New(commons.AliCloudErrorMessageVPC)
	}else {
		switchRequest.VpcId = request.VpcId
	}
	if request.VSwitchName == commons.ValueEmpty {
		switchRequest.VSwitchName = commons.AliCloudVSwitchName
	}else {
		switchRequest.VSwitchName = request.VSwitchName
	}
	if request.ZoneId == commons.ValueEmpty {
		switchRequest.ZoneId = commons.AliCloudZoneId
	}else {
		switchRequest.ZoneId = request.ZoneId
	}
	if request.CidrBlock == commons.ValueEmpty {
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
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	securityGroupRequest := ecs.CreateCreateSecurityGroupRequest()
	if request.VpcId == commons.ValueEmpty {

	}else {
		securityGroupRequest.VpcId = request.VpcId
	}
	if request.SecurityGroupName == commons.ValueEmpty {
		securityGroupRequest.SecurityGroupName = helpers.GenerateSecurityGroupName()
	}else {
		securityGroupRequest.SecurityGroupName = request.SecurityGroupName
	}
	securityGroupRequest.Description = commons.AliCloudSecurityGroupDescription
	securityGroupResponse, err := client.CreateSecurityGroup(securityGroupRequest)

	if err != nil {
		return nil, err
	}else {
		securityGroupId := securityGroupResponse.SecurityGroupId
		authorizeSecurityGroupRequest := ecs.CreateAuthorizeSecurityGroupRequest()
		authorizeSecurityGroupRequest.SecurityGroupId = securityGroupId
		authorizeSecurityGroupRequest.Description = commons.AliCloudSecurityGroupRuleDescription
		authorizeSecurityGroupRequest.NicType = "intranet"
		authorizeSecurityGroupRequest.IpProtocol = "tcp"
		authorizeSecurityGroupRequest.Policy = "accept"
		authorizeSecurityGroupRequest.PortRange = "22/22"
		authorizeSecurityGroupRequest.Priority = "1"
		authorizeSecurityGroupRequest.SourceCidrIp = "0.0.0.0/0"
		authorizeSecurityGroupRequest.Description = "This rule is created by BaaS"
		ports := [6]string{"22/22","7050/7054","2181/2181","2888/2888","3888/3888","9092/9093"}
		for _, port := range ports {
			authorizeSecurityGroupRequest.PortRange = port
			client.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
		}
		return &commons.CreateSecurityGroupResponse{SecurityGroupId:securityGroupId}, nil
	}
}

func (provider *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	runInstancesRequest := ecs.CreateRunInstancesRequest()
	helpers.TransferValuesBetweenRequest(runInstancesRequest,request)
	runInstancesResponse, err := client.RunInstances(runInstancesRequest)
	time.Sleep(time.Duration(5) * time.Second)
	if err == nil {
		instanceIdSet := runInstancesResponse.InstanceIdSets.InstanceIdSet
		instances := make(map[string]map[string]string)
		for _, instanceId := range instanceIdSet {
			mapData := make(map[string]string)
			instanceRequest := ecs.CreateDescribeInstanceAttributeRequest()
			instanceRequest.InstanceId = instanceId
			instanceResponse, err := client.DescribeInstanceAttribute(instanceRequest)
			if err == nil {
				if len(instanceResponse.PublicIpAddress.IpAddress) > 0 {
					mapData["public-ip"] = instanceResponse.PublicIpAddress.IpAddress[0]
				}
				mapData["instance-name"] = instanceResponse.InstanceName
			}
			instances[instanceId] = mapData
		}
		return &commons.RunInstancesResponse{Instances:instances}, nil
	} else {
		return nil, err
	}
}

func (provider *CloudProvider) DeleteInstance(instanceId string) {
	client, err := provider.GetClient()
	if err == nil {
		request := ecs.CreateDeleteInstanceRequest()
		request.InstanceId = instanceId
		client.DeleteInstance(request)
	}
}

func (provider *CloudProvider) GetZoneId(regionId,instanceChargeType string ) string {
	client, err := provider.GetClient()
	if err == nil {
		describeZonesRequest := ecs.CreateDescribeZonesRequest()
		describeZonesRequest.RegionId = regionId
		describeZonesRequest.InstanceChargeType = instanceChargeType
		describeZonesResponse, err := client.DescribeZones(describeZonesRequest)
		if err == nil {
			if len(describeZonesResponse.Zones.Zone) > 0 {
				return describeZonesResponse.Zones.Zone[0].ZoneId
			}else {
				return commons.AliCloudZoneId
			}
		}else {
			return commons.AliCloudZoneId
		}
	}else {
		return commons.AliCloudZoneId
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