package helpers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/commons"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)


func TransferValuesBetweenRequest(runInstancesRequest *ecs.RunInstancesRequest, request *commons.RunInstancesRequest){
	runInstancesRequest.ImageId = commons.AliCloudImageId
	runInstancesRequest.SystemDiskCategory = commons.AliCloudSystemDiskCategory
	runInstancesRequest.SystemDiskSize = commons.AliCloudSystemDiskSize
	runInstancesRequest.SecurityGroupId = request.SecurityGroupId
	runInstancesRequest.InstanceName = commons.AliCloudInstanceName
	runInstancesRequest.Description = commons.AliCloudInstanceDescription
	runInstancesRequest.InstanceType = commons.AliCloudInstanceType
	runInstancesRequest.InternetChargeType = commons.AliCloudInternetChargeType
	runInstancesRequest.InternetMaxBandwidthOut = requests.NewInteger(commons.AliCloudInternetMaxBandwidthOut)
	runInstancesRequest.HostName = commons.AliCloudInstanceHostName
	runInstancesRequest.Password = commons.AliCloudInstancePassword
	runInstancesRequest.InstanceChargeType = commons.AliCloudInstanceChargeType
	runInstancesRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	runInstancesRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)
	runInstancesRequest.PeriodUnit = "Month"
	runInstancesRequest.Period = requests.NewInteger(1)
	runInstancesRequest.IoOptimized = "optimized"
	runInstancesRequest.VSwitchId = request.VSwitchId
}