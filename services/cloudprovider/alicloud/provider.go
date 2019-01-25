package alicloud

import (
	"fmt"
	"time"
	"zig-cloud/commons"
	"zig-cloud/helpers"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/astaxie/beego"
)

type CloudProvider struct {
	Client *AliCloudClient
}

func (c *CloudProvider) CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error) {
	vpcRequest := BuildVpcRequest(request)

	raw, err := c.Client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
		return vpcClient.CreateVpc(vpcRequest)
	})
	if nil != err {
		beego.Error("CreateVpc failed:", err)
		return nil, helpers.ParseError(err)
	}
	vpcResponse, _ := raw.(*vpc.CreateVpcResponse)
	beego.Debug("CreateVpc successful vpcId =", vpcResponse.VpcId)

	err = c.WaitForVpc(vpcRequest.RegionId, vpcResponse.VpcId, commons.Available, 60)
	if err != nil {
		beego.Error("Timeout when WaitForVpc Available,", err)
		return nil, err
	}
	return &commons.CreateVpcResponse{VpcId: vpcResponse.VpcId}, nil
}

func (c *CloudProvider) DescribeVpc(regionId, vpcId string) (*vpc.DescribeVpcAttributeResponse, error) {
	vpcAttributeRequest := vpc.CreateDescribeVpcAttributeRequest()
	vpcAttributeRequest.RegionId = regionId
	vpcAttributeRequest.VpcId = vpcId

	raw, err := c.Client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
		return vpcClient.DescribeVpcAttribute(vpcAttributeRequest)
	})
	if nil != err {
		beego.Error("DescribeVpcAttribute failed:", err)
		return nil, helpers.ParseError(err)
	}
	vpcAttributeResponse, _ := raw.(*vpc.DescribeVpcAttributeResponse)
	beego.Debug("DescribeVpcAttribute successful:", vpcAttributeResponse)
	return vpcAttributeResponse, nil
}

func (c *CloudProvider) WaitForVpc(regionId, vpcId string, status commons.Status, timeout int) error {
	if timeout <= 0 {
		timeout = commons.DefaultTimeout
	}

	for {
		vpc, err := c.DescribeVpc(regionId, vpcId)
		if err != nil {
			return err
		}
		if vpc.Status == string(status) {
			break
		}
		timeout = timeout - commons.DefaultIntervalShort
		if timeout <= 0 {
			return fmt.Errorf(helpers.GetTimeoutMessage(vpcId, string(status)))
		}
		time.Sleep(commons.DefaultIntervalShort * time.Second)
	}
	return nil
}

func (c *CloudProvider) CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error) {
	vswitchRequest := BuildVSwitchRequest(request)

	raw, err := c.Client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
		return vpcClient.CreateVSwitch(vswitchRequest)
	})
	if nil != err {
		beego.Error("CreateVSwitch failed:", err)
		return nil, helpers.ParseError(err)
	}
	vswitchResponse, _ := raw.(*vpc.CreateVSwitchResponse)
	beego.Debug("CreateVSwitch successful VSwitchId =", vswitchResponse.VSwitchId)
	return &commons.CreateVSwitchResponse{VSwitchId: vswitchResponse.VSwitchId}, nil
}

func (c *CloudProvider) CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error) {
	securityGroupRequest := BuildSecurityGroupRequest(request)

	raw, err := c.Client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.CreateSecurityGroup(securityGroupRequest)
	})
	if nil != err {
		beego.Error("CreateSecurityGroup failed:", err)
		return nil, helpers.ParseError(err)
	}
	securityGroupResponse, _ := raw.(*ecs.CreateSecurityGroupResponse)
	if nil == securityGroupResponse {
		return nil, fmt.Errorf("Creating security group got a nil response")
	}
	beego.Debug("CreateSecurityGroup successful SecurityGroupId =", securityGroupResponse.SecurityGroupId)
	return &commons.CreateSecurityGroupResponse{SecurityGroupId: securityGroupResponse.SecurityGroupId}, nil
}

func (c *CloudProvider) AuthorizeSecurityGroup(request *commons.AuthorizeSecurityGroupRequest) (*commons.AuthorizeSecurityGroupResponse, error) {
	authorizeSecurityGroupRequest := BuildAuthorizeSecurityGroupRequest(request)

	raw, err := c.Client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
	})
	if nil != err {
		beego.Error("AuthorizeSecurityGroup failed:", err)
		return nil, helpers.ParseError(err)
	}
	authorizeSecurityGroupResponse, _ := raw.(*ecs.AuthorizeSecurityGroupResponse)
	beego.Debug("AuthorizeSecurityGroup successful RequestId =", authorizeSecurityGroupResponse.RequestId)
	return &commons.AuthorizeSecurityGroupResponse{RequestId: authorizeSecurityGroupResponse.RequestId}, nil
}

func (c *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	//runInstancesRequest := BuildInstanceRequest(request)
	//runInstancesResponse, err := client.RunInstances(runInstancesRequest)
	//time.Sleep(time.Duration(5) * time.Second)
	//if err == nil {
	//	instanceIdSet := runInstancesResponse.InstanceIdSets.InstanceIdSet
	//	instances := make(map[string]map[string]string)
	//	for _, instanceId := range instanceIdSet {
	//		mapData := make(map[string]string)
	//		instanceRequest := ecs.CreateDescribeInstanceAttributeRequest()
	//		instanceRequest.InstanceId = instanceId
	//		instanceResponse, err := client.DescribeInstanceAttribute(instanceRequest)
	//		if err == nil {
	//			if len(instanceResponse.PublicIpAddress.IpAddress) > 0 {
	//				mapData["public-ip"] = instanceResponse.PublicIpAddress.IpAddress[0]
	//			}
	//			mapData["instance-name"] = instanceResponse.InstanceName
	//		}
	//		instances[instanceId] = mapData
	//	}
	//	return &commons.RunInstancesResponse{Instances: instances}, nil
	//} else {
	//	return nil, err
	//}
	return nil, nil
}

func (c *CloudProvider) DeleteInstance(instanceId string) (*commons.DeleteInstanceResponse, error) {
	deleteInstanceRequest := ecs.CreateDeleteInstanceRequest()
	deleteInstanceRequest.InstanceId = instanceId

	raw, err := c.Client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DeleteInstance(deleteInstanceRequest)
	})
	if nil != err {
		beego.Error("DeleteInstance failed:", err)
		return nil, helpers.ParseError(err)
	}
	deleteInstanceResponse, _ := raw.(*ecs.DeleteInstanceResponse)
	beego.Debug("DeleteInstance successful RequestId =", deleteInstanceResponse.RequestId)
	return &commons.DeleteInstanceResponse{RequestId: deleteInstanceResponse.RequestId}, nil
}

func (c *CloudProvider) DescribeRegions() (*ecs.DescribeRegionsResponse, error) {
	describeRegionsRequest := ecs.CreateDescribeRegionsRequest()
	raw, err := c.Client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeRegions(describeRegionsRequest)
	})
	if nil != err {
		beego.Error("DescribeRegions failed:", err)
		return nil, helpers.ParseError(err)
	}
	describeRegionsResponse, _ := raw.(*ecs.DescribeRegionsResponse)
	beego.Debug("DescribeRegions successful:", describeRegionsResponse)
	return describeRegionsResponse, nil
}

func (c *CloudProvider) DescribeZones(regionId string) ([]*commons.DescribeZonesResponse, error) {
	var describeZonesSlice []*commons.DescribeZonesResponse
	describeZonesRequest := ecs.CreateDescribeZonesRequest()
	describeZonesRequest.RegionId = regionId

	raw, err := c.Client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeZones(describeZonesRequest)
	})
	if nil != err {
		beego.Error("DescribeZones failed:", err)
		return nil, helpers.ParseError(err)
	}
	describeZonesResponse, _ := raw.(*ecs.DescribeZonesResponse)
	beego.Debug("DescribeZones successful:", describeZonesResponse)
	if len(describeZonesResponse.Zones.Zone) > 0 {
		for _, v := range describeZonesResponse.Zones.Zone {
			describeZonesSlice = append(describeZonesSlice, &commons.DescribeZonesResponse{ZoneId: v.ZoneId, LocalName: v.LocalName})
		}
		return describeZonesSlice, nil
	}
	return describeZonesSlice, nil
}
