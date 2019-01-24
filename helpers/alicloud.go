package helpers

import (
	"zig-cloud/commons"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func TransferValuesBetweenRequest(runInstancesRequest *ecs.RunInstancesRequest, request *commons.RunInstancesRequest) {
	if request.ImageId == commons.ValueEmpty {
		runInstancesRequest.ImageId = commons.AliCloudImageId
	} else {
		runInstancesRequest.ImageId = request.ImageId
	}
	runInstancesRequest.SystemDiskCategory = commons.AliCloudSystemDiskCategory
	runInstancesRequest.SystemDiskSize = commons.AliCloudSystemDiskSize
	runInstancesRequest.SecurityGroupId = request.SecurityGroupId
	runInstancesRequest.InstanceName = commons.AliCloudInstanceName
	runInstancesRequest.Description = commons.AliCloudInstanceDescription
	if request.InstanceType == commons.ValueEmpty {
		runInstancesRequest.InstanceType = commons.AliCloudInstanceType
	} else {
		runInstancesRequest.InstanceType = request.InstanceType
	}
	runInstancesRequest.InternetChargeType = commons.AliCloudInternetChargeType
	runInstancesRequest.InternetMaxBandwidthOut = requests.NewInteger(commons.AliCloudInternetMaxBandwidthOut)
	runInstancesRequest.HostName = commons.AliCloudInstanceHostName
	runInstancesRequest.Password = commons.AliCloudInstancePassword
	if request.InstanceChargeType == commons.ValueEmpty {
		runInstancesRequest.InstanceChargeType = commons.AliCloudInstanceChargeTypePostPaid
	} else {
		runInstancesRequest.InstanceChargeType = request.InstanceChargeType
	}
	if runInstancesRequest.InstanceChargeType == commons.AliCloudInstanceChargeTypePrePaid {
		runInstancesRequest.PeriodUnit = commons.AliCloudInstanceChargeTypePeriodUnit
		runInstancesRequest.Period = requests.NewInteger(commons.AliCloudInstanceChargeTypePeriod)
	}
	if request.Amount == commons.ValueEmpty {
		runInstancesRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	} else {
		runInstancesRequest.Amount = requests.Integer(request.Amount)
	}
	runInstancesRequest.IoOptimized = commons.AliCloudIoOptimized
	runInstancesRequest.VSwitchId = request.VSwitchId
	runInstancesRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)
}
