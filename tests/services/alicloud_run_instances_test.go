package services

import (
	"testing"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"zig-cloud/commons"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func TestAliCloudRunInstances(t *testing.T) {
	client := GetClient()

	runInstanceRequest := ecs.CreateRunInstancesRequest()

	runInstanceRequest.ImageId = commons.AliCloudImageId
	runInstanceRequest.SystemDiskCategory = commons.AliCloudSystemDiskCategory
	runInstanceRequest.SystemDiskSize = commons.AliCloudSystemDiskSize
	//runInstanceRequest.SecurityGroupId = commons.AliCloudSecurityGroupId
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

	response, err := client.RunInstances(runInstanceRequest)

	if err == nil {
		fmt.Println(response.GetHttpContentString())
	}else {
		t.Error(err)
	}
}
