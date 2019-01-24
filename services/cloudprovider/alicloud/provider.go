package alicloud

import (
	"errors"
	"time"
	"zig-cloud/commons"
	"zig-cloud/helpers"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/astaxie/beego"
)

type CloudProvider struct {
	Client *AliCloudClient
}

func (provider *CloudProvider) CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	vpcRequest := ecs.CreateCreateVpcRequest()
	if request.VpcName == commons.ValueEmpty {
		vpcRequest.VpcName = commons.AliCloudVpcName
	} else {
		vpcRequest.VpcName = request.VpcName
	}
	if request.RegionId == commons.ValueEmpty {
		vpcRequest.RegionId = commons.AliCloudRegionId
	} else {
		vpcRequest.RegionId = request.RegionId
	}
	if request.CidrBlock == commons.ValueEmpty {
		vpcRequest.CidrBlock = commons.AliCloudVpcCidrBlock
	} else {
		vpcRequest.CidrBlock = request.CidrBlock
	}
	vpcRequest.Description = commons.AliCloudVpcDescription
	vpcRequest.ClientToken = utils.GetUUIDV4()
	vpcResponse, err := client.CreateVpc(vpcRequest)
	if err == nil {
		beego.Debug("CreateVpc successful vpcId =", vpcResponse.VpcId)
		return &commons.CreateVpcResponse{VpcId: vpcResponse.VpcId}, nil
	} else {
		beego.Error("CreateVpc failed:", err)
		return nil, helpers.ParseError(err)
	}
}

func (provider *CloudProvider) CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	switchRequest := ecs.CreateCreateVSwitchRequest()
	if request.VSwitchName == commons.ValueEmpty {
		switchRequest.VSwitchName = commons.AliCloudVSwitchName
	} else {
		switchRequest.VSwitchName = request.VSwitchName
	}
	if request.VpcId == commons.ValueEmpty {
		return nil, errors.New(commons.AliCloudErrorMessageVPC)
	} else {
		switchRequest.VpcId = request.VpcId
	}
	if request.ZoneId == commons.ValueEmpty {
		switchRequest.ZoneId = commons.AliCloudZoneId
	} else {
		switchRequest.ZoneId = request.ZoneId
	}
	if request.RegionId == commons.ValueEmpty {
		switchRequest.RegionId = commons.AliCloudRegionId
	} else {
		switchRequest.RegionId = request.RegionId
	}
	if request.CidrBlock == commons.ValueEmpty {
		switchRequest.CidrBlock = commons.AliCloudSwitchCidrBlock
	} else {
		switchRequest.CidrBlock = request.CidrBlock
	}
	switchRequest.Description = commons.AliCloudVSwitchDescription
	switchRequest.ClientToken = utils.GetUUIDV4()
	switchResponse, err := client.CreateVSwitch(switchRequest)
	if err == nil {
		beego.Debug("CreateVSwitch successful VSwitchId =", switchResponse.VSwitchId)
		return &commons.CreateVSwitchResponse{VSwitchId: switchResponse.VSwitchId}, nil
	} else {
		beego.Error("CreateVSwitch failed:", err)
		return nil, helpers.ParseError(err)
	}
}

func (provider *CloudProvider) CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	securityGroupRequest := ecs.CreateCreateSecurityGroupRequest()
	securityGroupRequest.RegionId = request.RegionId
	if request.VpcId == commons.ValueEmpty {

	} else {
		securityGroupRequest.VpcId = request.VpcId
	}
	if request.SecurityGroupName == commons.ValueEmpty {
		securityGroupRequest.SecurityGroupName = helpers.GenerateSecurityGroupName()
	} else {
		securityGroupRequest.SecurityGroupName = request.SecurityGroupName
	}
	securityGroupRequest.Description = commons.AliCloudSecurityGroupDescription
	securityGroupRequest.ClientToken = utils.GetUUIDV4()
	securityGroupResponse, err := client.CreateSecurityGroup(securityGroupRequest)
	if err != nil {
		return nil, helpers.ParseError(err)
	}
	return &commons.CreateSecurityGroupResponse{SecurityGroupId: securityGroupResponse.SecurityGroupId}, nil
}

func (provider *CloudProvider) CreateAuthorizeSecurityGroup(request *commons.AuthorizeSecurityGroupRequest) (*commons.AuthorizeSecurityGroupResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}

	authorizeSecurityGroupRequest := ecs.CreateAuthorizeSecurityGroupRequest()
	authorizeSecurityGroupRequest.RegionId = request.RegionId
	authorizeSecurityGroupRequest.SecurityGroupId = request.SecurityGroupId
	authorizeSecurityGroupRequest.IpProtocol = "tcp"
	authorizeSecurityGroupRequest.PortRange = request.PortRange
	authorizeSecurityGroupRequest.NicType = "intranet"
	authorizeSecurityGroupRequest.SourceCidrIp = "0.0.0.0/0"
	authorizeSecurityGroupRequest.Policy = "accept"
	authorizeSecurityGroupRequest.Priority = "1"
	authorizeSecurityGroupRequest.Description = commons.AliCloudSecurityGroupRuleDescription
	authorizeSecurityGroupRequest.ClientToken = utils.GetUUIDV4()
	authorizeSecurityGroupResponse, err := client.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
	if nil != err {
		return nil, helpers.ParseError(err)
	}
	return &commons.AuthorizeSecurityGroupResponse{RequestId: authorizeSecurityGroupResponse.RequestId}, nil
}

func (provider *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}

	runInstancesRequest := ecs.CreateRunInstancesRequest()
	helpers.TransferValuesBetweenRequest(runInstancesRequest, request)
	runInstancesRequest.ClientToken = utils.GetUUIDV4()
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
		return &commons.RunInstancesResponse{Instances: instances}, nil
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

func (provider *CloudProvider) DescribeRegions() (*ecs.DescribeRegionsResponse, error) {
	client, err := provider.GetClient()
	if err != nil {
		return nil, err
	}
	describeRegionsRequest := ecs.CreateDescribeRegionsRequest()
	return client.DescribeRegions(describeRegionsRequest)
}

func (provider *CloudProvider) DescribeZones(regionId string) ([]*commons.DescribeZonesResponse, error) {
	var describeZonesSlice []*commons.DescribeZonesResponse
	client, err := provider.GetClient()
	if err != nil {
		return describeZonesSlice, err
	}
	describeZonesRequest := ecs.CreateDescribeZonesRequest()
	describeZonesRequest.RegionId = regionId
	describeZonesResponse, err := client.DescribeZones(describeZonesRequest)
	if err == nil {
		if len(describeZonesResponse.Zones.Zone) > 0 {
			for _, v := range describeZonesResponse.Zones.Zone {
				describeZonesSlice = append(describeZonesSlice, &commons.DescribeZonesResponse{ZoneId: v.ZoneId, LocalName: v.LocalName})
			}
			return describeZonesSlice, nil
		}
	}
	return describeZonesSlice, helpers.ParseError(err)
}

//
//func (provider *CloudProvider) DescribeVpc(vpcId string) ([]*commons.DescribeZonesResponse, error) {
//	client, err := provider.GetClient()
//	if err != nil {
//		return describeZonesSlice, err
//	}
//	describeZonesRequest := vpc.CreateDescribeVpcAttributeRequest()
//	describeZonesRequest.RegionId = regionId
//	describeZonesResponse, err := client.DescribeVpcAttribute(describeZonesRequest)
//	if err == nil {
//		if len(describeZonesResponse.Zones.Zone) > 0 {
//			for _, v := range describeZonesResponse.Zones.Zone {
//				describeZonesSlice = append(describeZonesSlice, &commons.DescribeZonesResponse{ZoneId: v.ZoneId, LocalName: v.LocalName})
//			}
//			return describeZonesSlice, nil
//		}
//	}
//	return describeZonesSlice, helpers.ParseError(err)
//}
//
//func (provider *CloudProvider) WaitForVpc(vpcId string, status Status, timeout int) error {
//	if timeout <= 0 {
//		timeout = commons.DefaultTimeout
//	}
//
//	for {
//		vpc, err := provider.DescribeVpc(vpcId)
//		if err != nil {
//			return err
//		}
//		if vpc.Status == string(status) {
//			break
//		}
//		timeout = timeout - DefaultIntervalShort
//		if timeout <= 0 {
//			return GetTimeErrorFromString(GetTimeoutMessage("VPC", string(status)))
//		}
//		time.Sleep(DefaultIntervalShort * time.Second)
//	}
//	return nil
//}

func (provider *CloudProvider) GetClient() (*ecs.Client, error) {
	client, err := ecs.NewClientWithAccessKey(provider.Client.RegionId, provider.Client.AccessKeyId, provider.Client.AccessKeySecret)
	if err == nil {
		return client, nil
	} else {
		beego.Error(err)
		return nil, err
	}
}
