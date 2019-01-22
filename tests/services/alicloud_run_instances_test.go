package services

import (
	"testing"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"zig-cloud/commons"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"zig-cloud/helpers"
	"time"
)

func TestAliCloudDescribeZones(t *testing.T) {
	client := GetClient()
	describeZonesRequest := ecs.CreateDescribeZonesRequest()
	describeZonesRequest.RegionId = "cn-huhehaote"
	describeZonesRequest.InstanceChargeType = "PrePaid"
	describeZonesResponse, err := client.DescribeZones(describeZonesRequest)
	if err == nil {
		fmt.Println(">>> no errors")
		fmt.Println(describeZonesResponse.GetHttpContentString())
		//fmt.Println(describeZonesResponse.Zones)
		zoneId := describeZonesResponse.Zones.Zone[0].ZoneId
		fmt.Println(">>>> zone id " + zoneId)
	} else {
		fmt.Println(">>> errors happened")
		t.Error(err)
	}
}

func TestAliCloudCreateSwitch(t *testing.T) {
	client := GetClient()

	vpcRequest := ecs.CreateCreateVpcRequest()
	vpcRequest.VpcName = "BaaSVPC"
	vpcRequest.CidrBlock = "172.16.0.0/12"

	vpcResponse, err := client.CreateVpc(vpcRequest)

	if err == nil {
		fmt.Println(vpcResponse.GetHttpContentString())
	} else {
		t.Error(err)
		fmt.Println(err)
	}

	time.Sleep(time.Duration(5) * time.Second)

	switchRequest := ecs.CreateCreateVSwitchRequest()
	switchRequest.VpcId = vpcResponse.VpcId
	switchRequest.CidrBlock = "172.16.0.0/24"
	switchRequest.ZoneId = "cn-huhehaote-a"
	switchResponse, err := client.CreateVSwitch(switchRequest)

	if err == nil {
		fmt.Println(">>> no errors")
		fmt.Println(switchResponse.GetHttpContentString())
	} else {
		fmt.Println(">>> errors happened")
		t.Error(err)
		fmt.Println(err)
	}
}

func TestAliCloudRunInstances(t *testing.T) {
	client := GetClient()

	vpcRequest := ecs.CreateCreateVpcRequest()
	vpcRequest.VpcName = "BaaSVPC"
	vpcRequest.CidrBlock = "172.16.0.0/12"

	vpcResponse, err := client.CreateVpc(vpcRequest)

	if err == nil {
		fmt.Println(vpcResponse.GetHttpContentString())
	}

	time.Sleep(time.Duration(5) * time.Second)

	describeZonesRequest := ecs.CreateDescribeZonesRequest()
	describeZonesRequest.RegionId = "cn-huhehaote"
	describeZonesRequest.InstanceChargeType = "PrePaid"

	describeZonesResponse, err := client.DescribeZones(describeZonesRequest)

	if err == nil {
		//fmt.Println(describeZonesResponse.GetHttpContentString())
		zoneId := describeZonesResponse.Zones.Zone[0].ZoneId
		fmt.Println(">>>> zone id " + zoneId)
	}else {
		fmt.Println(err)
	}

	switchRequest := ecs.CreateCreateVSwitchRequest()
	switchRequest.VpcId = vpcResponse.VpcId
	switchRequest.CidrBlock = "172.16.0.0/24"
	switchRequest.ZoneId = "cn-huhehaote-a"

	switchResponse, err := client.CreateVSwitch(switchRequest)

	if err == nil {
		fmt.Println(switchResponse.GetHttpContentString())
	}else{
		fmt.Println(err)
	}

	time.Sleep(time.Duration(5) * time.Second)

	request := ecs.CreateCreateSecurityGroupRequest()
	request.SecurityGroupName = helpers.GenerateSecurityGroupName()
	request.Description = commons.AliCloudSecurityGroupDescription
	request.VpcId = vpcResponse.VpcId
	response, err := client.CreateSecurityGroup(request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	fmt.Println("The security group ID is " + response.SecurityGroupId)
	fmt.Println(response.GetHttpContentString())

	authorizeSecurityGroupRequest := ecs.CreateAuthorizeSecurityGroupRequest();
	authorizeSecurityGroupRequest.SecurityGroupId = response.SecurityGroupId
	authorizeSecurityGroupRequest.NicType = "intranet"
	authorizeSecurityGroupRequest.IpProtocol = "tcp"
	authorizeSecurityGroupRequest.Policy = "accept"
	authorizeSecurityGroupRequest.PortRange = "22/22"
	authorizeSecurityGroupRequest.Priority = "1"
	authorizeSecurityGroupRequest.SourceCidrIp = "0.0.0.0/0"
	authorizeSecurityGroupRequest.Description = "This rule is created by BaaS"
	client.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)

	runInstanceRequest := ecs.CreateRunInstancesRequest()

	runInstanceRequest.ImageId = commons.AliCloudImageId
	runInstanceRequest.SystemDiskCategory = commons.AliCloudSystemDiskCategory
	runInstanceRequest.SystemDiskSize = commons.AliCloudSystemDiskSize
	runInstanceRequest.SecurityGroupId = response.SecurityGroupId
	runInstanceRequest.InstanceName = commons.AliCloudInstanceName
	runInstanceRequest.Description = commons.AliCloudInstanceDescription
	runInstanceRequest.InstanceType = commons.AliCloudInstanceType
	runInstanceRequest.InternetChargeType = commons.AliCloudInternetChargeType
	runInstanceRequest.InternetMaxBandwidthOut = requests.NewInteger(commons.AliCloudInternetMaxBandwidthOut)
	runInstanceRequest.HostName = commons.AliCloudInstanceHostName
	runInstanceRequest.Password = commons.AliCloudInstancePassword
	runInstanceRequest.InstanceChargeType = commons.AliCloudInstanceChargeTypePostPaid
	runInstanceRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	runInstanceRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)
	runInstanceRequest.PeriodUnit = "Month"
	runInstanceRequest.Period = requests.NewInteger(1)

	runInstanceRequest.IoOptimized = "optimized"

	runInstanceRequest.VSwitchId = switchResponse.VSwitchId

	instanceResponse, err := client.RunInstances(runInstanceRequest)

	if err == nil {
		fmt.Println(instanceResponse.GetHttpContentString())
	}else {
		t.Error(err)
	}
}
