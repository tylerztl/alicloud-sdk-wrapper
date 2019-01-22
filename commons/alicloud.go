package commons


const (
	SeparatorHype string = "-"
	ValueEmpty = ""
)

const (

	AliCloudRegionId string = "cn-huhehaote"
	AliCloudZoneId string = "cn-huhehaote-a"

	// VPC
	AliCloudVPCName string = "baas-vpc"
	AliCloudVPCCidrBlock string = "172.16.0.0/12"

	// virtual switch
	AliCloudVSwitchName string = "baas-vswitch"
	AliCloudSwitchCidrBlock string = "172.16.0.0/24"

	// security group
	AliCloudSecurityGroupName string = "bass-sg"
	AliCloudSecurityGroupDescription string = "It is applied to BaaS instances"
	AliCloudSecurityGroupRuleName string = "baas-sg-rule"
	AliCloudSecurityGroupRuleDescription string = "baas and blockchain network"

	// instance
	AliCloudInstanceType string = "ecs.c5.large"
	AliCloudImageId string = "ubuntu_16_0402_64_20G_alibase_20180409.vhd"
	AliCloudSystemDiskCategory string = "cloud_efficiency"
	AliCloudSystemDiskSize string = "40"
	AliCloudSecurityGroupId string = "system"
	AliCloudInstanceName string = "bass-instance"
	AliCloudInstanceDescription string = "This instance is created by BaaS"
	AliCloudInternetChargeType string = "PayByBandwidth"
	AliCloudInternetMaxBandwidthOut int = 1
	AliCloudInstanceHostName string = "instance-baas"
	AliCloudInstancePassword string = "Pass@w0rd"
	AliCloudInstanceChargeTypePrePaid string = "PrePaid"
	AliCloudInstanceChargeTypePostPaid string = "PostPaid"
	AliCloudInstanceChargeTypePeriodUnit string = "Month"
	AliCloudInstanceChargeTypePeriod int = 1
	AliCloudInstanceAmount int = 1
	AliCloudIoOptimized string = "optimized"
	AliCloudDryRun bool = false

	// error messages
	AliCloudErrorMessageVPC string = "no vpc"
	AliCloudErrorMessageInvalidData string = "invalid data"

)
