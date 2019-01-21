package commons

// define the constant variables here

const (
	SeparatorHype string = "-"
)

const (

	// security group
	AliCloudSecurityGroupName string = "bass-sg"
	AliCloudSecurityGroupDescription string = "It is applied to BaaS instances"
	AliCloudSecurityGroupRuleName string = "baas-sg-rule"

	// instance
	AliCloudInstanceType string = "ecs.c5.large"
	AliCloudImageId string = "ubuntu_16_0402_64_20G_alibase_20180409.vhd"
	AliCloudSystemDiskCategory string = "cloud_efficiency"
	AliCloudSystemDiskSize string = "40"
	AliCloudSecurityGroupId string = ""
	AliCloudInstanceName string = "bass-instance"
	AliCloudInstanceDescription string = "This instance is created by BaaS"
	AliCloudInternetChargeType string = "PayByBandwidth"
	AliCloudInternetMaxBandwidthOut int = 100
	AliCloudInstanceHostName string = "instance-baas"
	AliCloudInstancePassword string = "Pass@w0rd"
	AliCloudInstanceChargeType string = "PrePaid"
	AliCloudInstanceAmount int = 1
	AliCloudDryRun bool = true
)
