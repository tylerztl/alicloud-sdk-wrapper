package helpers

import (
	"zig-cloud/commons"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func TransferValuesBetweenRequest(runInstancesRequest *ecs.RunInstancesRequest, request *commons.RunInstancesRequest) {
	runInstancesRequest.RegionId = request.RegionId
	runInstancesRequest.ZoneId = request.ZoneId
	if request.ImageId == commons.ValueEmpty {
		runInstancesRequest.ImageId = commons.AliCloudImageId
	} else {
		runInstancesRequest.ImageId = request.ImageId
	}
	if request.InstanceType == commons.ValueEmpty {
		runInstancesRequest.InstanceType = commons.AliCloudInstanceType
	} else {
		runInstancesRequest.InstanceType = request.InstanceType
	}
	runInstancesRequest.SecurityGroupId = request.SecurityGroupId
	runInstancesRequest.VSwitchId = request.VSwitchId
	if request.InstanceChargeType == commons.ValueEmpty {
		runInstancesRequest.InstanceChargeType = commons.AliCloudInstanceChargeTypePostPaid
	} else {
		runInstancesRequest.InstanceChargeType = request.InstanceChargeType
	}
	if runInstancesRequest.InstanceChargeType == commons.AliCloudInstanceChargeTypePrePaid {
		runInstancesRequest.PeriodUnit = commons.AliCloudInstanceChargeTypePeriodUnit
		runInstancesRequest.Period = requests.NewInteger(commons.AliCloudInstanceChargeTypePeriod)
	}
	runInstancesRequest.HostName = commons.AliCloudInstanceHostName
	if request.Password == commons.ValueEmpty {
		runInstancesRequest.Password = commons.AliCloudInstancePassword
	} else {
		runInstancesRequest.Password = request.Password
	}

	runInstancesRequest.InternetChargeType = commons.AliCloudInternetChargeType
	runInstancesRequest.IoOptimized = commons.AliCloudIoOptimized

	runInstancesRequest.SystemDiskCategory = commons.AliCloudSystemDiskCategory
	runInstancesRequest.SystemDiskSize = commons.AliCloudSystemDiskSize

	runInstancesRequest.InternetMaxBandwidthOut = requests.NewInteger(commons.AliCloudInternetMaxBandwidthOut)
	runInstancesRequest.InstanceName = commons.AliCloudInstanceName
	runInstancesRequest.Description = commons.AliCloudInstanceDescription

	if request.Amount == commons.ValueEmpty {
		runInstancesRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	} else {
		runInstancesRequest.Amount = requests.Integer(request.Amount)
	}
	runInstancesRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)
}
