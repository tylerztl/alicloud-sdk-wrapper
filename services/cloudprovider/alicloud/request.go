package alicloud

import (
	"zig-cloud/commons"
	"zig-cloud/helpers"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

func BuildVpcRequest(request *commons.CreateVpcRequest) *vpc.CreateVpcRequest {
	vpcRequest := vpc.CreateCreateVpcRequest()

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

	return vpcRequest
}

func BuildVSwitchRequest(request *commons.CreateVSwitchRequest) *vpc.CreateVSwitchRequest {
	vswitchRequest := vpc.CreateCreateVSwitchRequest()

	if request.VSwitchName == commons.ValueEmpty {
		vswitchRequest.VSwitchName = commons.AliCloudVSwitchName
	} else {
		vswitchRequest.VSwitchName = request.VSwitchName
	}
	vswitchRequest.VpcId = request.VpcId
	if request.ZoneId == commons.ValueEmpty {
		vswitchRequest.ZoneId = commons.AliCloudZoneId
	} else {
		vswitchRequest.ZoneId = request.ZoneId
	}
	if request.RegionId == commons.ValueEmpty {
		vswitchRequest.RegionId = commons.AliCloudRegionId
	} else {
		vswitchRequest.RegionId = request.RegionId
	}
	if request.CidrBlock == commons.ValueEmpty {
		vswitchRequest.CidrBlock = commons.AliCloudSwitchCidrBlock
	} else {
		vswitchRequest.CidrBlock = request.CidrBlock
	}
	vswitchRequest.Description = commons.AliCloudVSwitchDescription
	vswitchRequest.ClientToken = utils.GetUUIDV4()

	return vswitchRequest
}

func BuildSecurityGroupRequest(request *commons.CreateSecurityGroupRequest) *ecs.CreateSecurityGroupRequest {
	securityGroupRequest := ecs.CreateCreateSecurityGroupRequest()

	if request.RegionId == commons.ValueEmpty {
		securityGroupRequest.RegionId = commons.AliCloudRegionId
	} else {
		securityGroupRequest.RegionId = request.RegionId
	}
	securityGroupRequest.VpcId = request.VpcId
	if request.SecurityGroupName == commons.ValueEmpty {
		securityGroupRequest.SecurityGroupName = helpers.GenerateSecurityGroupName()
	} else {
		securityGroupRequest.SecurityGroupName = request.SecurityGroupName
	}
	securityGroupRequest.Description = commons.AliCloudSecurityGroupDescription
	securityGroupRequest.ClientToken = utils.GetUUIDV4()

	return securityGroupRequest
}

func BuildAuthorizeSecurityGroupRequest(request *commons.AuthorizeSecurityGroupRequest) *ecs.AuthorizeSecurityGroupRequest {
	authorizeSecurityGroupRequest := ecs.CreateAuthorizeSecurityGroupRequest()

	if request.RegionId == commons.ValueEmpty {
		authorizeSecurityGroupRequest.RegionId = commons.AliCloudRegionId
	} else {
		authorizeSecurityGroupRequest.RegionId = request.RegionId
	}
	authorizeSecurityGroupRequest.SecurityGroupId = request.SecurityGroupId
	if request.IpProtocol == commons.ValueEmpty {
		authorizeSecurityGroupRequest.IpProtocol = commons.AliCloudSecurityGroupRuleIpProtocol
	} else {
		authorizeSecurityGroupRequest.IpProtocol = request.IpProtocol
	}
	authorizeSecurityGroupRequest.PortRange = request.PortRange
	authorizeSecurityGroupRequest.NicType = commons.AliCloudSecurityGroupRuleNicType
	authorizeSecurityGroupRequest.SourceCidrIp = commons.AliCloudSecurityGroupRuleSourceCidrIp
	authorizeSecurityGroupRequest.Policy = commons.AliCloudSecurityGroupRulePolicy
	authorizeSecurityGroupRequest.Priority = commons.AliCloudSecurityGroupRulePriority
	authorizeSecurityGroupRequest.Description = commons.AliCloudSecurityGroupRuleDescription
	authorizeSecurityGroupRequest.ClientToken = utils.GetUUIDV4()

	return authorizeSecurityGroupRequest
}

func BuildInstanceRequest(request *commons.RunInstancesRequest) *ecs.RunInstancesRequest {
	runInstancesRequest := ecs.CreateRunInstancesRequest()

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
	if request.HostName == commons.ValueEmpty {
		runInstancesRequest.HostName = commons.AliCloudInstanceHostName
	} else {
		runInstancesRequest.HostName = request.HostName
	}
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
	runInstancesRequest.UniqueSuffix = requests.NewBoolean(commons.AliCloudUniqueSuffix)
	runInstancesRequest.Tag = &[]ecs.RunInstancesTag{{Key: request.TagKey, Value: request.TagValue}}
	runInstancesRequest.Description = commons.AliCloudInstanceDescription

	if request.Amount <= 0 {
		runInstancesRequest.Amount = requests.NewInteger(commons.AliCloudInstanceAmount)
	} else {
		runInstancesRequest.Amount = requests.NewInteger(request.Amount)
	}
	runInstancesRequest.DryRun = requests.NewBoolean(commons.AliCloudDryRun)
	runInstancesRequest.ClientToken = utils.GetUUIDV4()

	return runInstancesRequest
}
