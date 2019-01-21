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

func TestAliCloudRunInstances(t *testing.T) {
	client := GetClient()

	vpcRequest := ecs.CreateCreateVpcRequest()
	vpcRequest.VpcName = "BaaSVPC"
	vpcRequest.CidrBlock = "172.16.0.0/12"

	vpcResponse, err := client.CreateVpc(vpcRequest)

	if err == nil {
		fmt.Println(vpcResponse.GetHttpContentString())
	}

	time.Sleep(time.Duration(10) * time.Second)

	switchRequest := ecs.CreateCreateVSwitchRequest()
	switchRequest.VpcId = vpcResponse.VpcId
	switchRequest.CidrBlock = "172.16.0.0/24"

	switchResponse, err := client.CreateVSwitch(switchRequest)

	if err == nil {
		fmt.Println(switchResponse.GetHttpContentString())
	}

	time.Sleep(time.Duration(10) * time.Second)

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
	runInstanceRequest.InstanceChargeType = commons.AliCloudInstanceChargeType
	runInstanceRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	runInstanceRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)

	runInstanceRequest.IoOptimized = "none"

	runInstanceRequest.VSwitchId = switchResponse.VSwitchId

	instanceResponse, err := client.RunInstances(runInstanceRequest)

	if err == nil {
		fmt.Println(instanceResponse.GetHttpContentString())
	}else {
		t.Error(err)
	}
}
