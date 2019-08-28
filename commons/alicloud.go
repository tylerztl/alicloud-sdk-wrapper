package commons

const (
	SeparatorHype string = "-"
	ValueEmpty    string = ""
)

const (
	DefaultTimeout        = 120
	DefaultIntervalShort  = 5
	DefaultTimeoutMedium  = 500
	DefaultIntervalMedium = 20
	DefaultLongTimeout    = 1000
	DefaultIntervalLong   = 30
)

type Status string

const (
	Pending     = Status("Pending")
	Creating    = Status("Creating")
	Running     = Status("Running")
	Available   = Status("Available")
	Unavailable = Status("Unavailable")
	Modifying   = Status("Modifying")
	Deleting    = Status("Deleting")
	Starting    = Status("Starting")
	Stopping    = Status("Stopping")
	Stopped     = Status("Stopped")
	Normal      = Status("Normal")
	Changing    = Status("Changing")
	Online      = Status("online")
	Configuring = Status("configuring")

	Associating   = Status("Associating")
	Unassociating = Status("Unassociating")
	InUse         = Status("InUse")
	DiskInUse     = Status("In_use")

	Active   = Status("Active")
	Inactive = Status("Inactive")
	Idle     = Status("Idle")

	SoldOut = Status("SoldOut")

	InService      = Status("InService")
	Removing       = Status("Removing")
	DisabledStatus = Status("Disabled")

	Init            = Status("Init")
	Provisioning    = Status("Provisioning")
	Updating        = Status("Updating")
	FinancialLocked = Status("FinancialLocked")

	PUBLISHED   = Status("Published")
	NOPUBLISHED = Status("NonPublished")
)

const (
	AliCloudRegionId string = "cn-huhehaote"
	AliCloudZoneId   string = "cn-huhehaote-a"

	// vpc
	AliCloudVpcName        string = "vpc-alicloud-sdk-wrapper"
	AliCloudVpcCidrBlock   string = "172.16.0.0/12"
	AliCloudVpcDescription string = "This vpc is created by alicloud-sdk-wrapper"

	// virtual switch
	AliCloudVSwitchName        string = "vsw-alicloud-sdk-wrapper"
	AliCloudSwitchCidrBlock    string = "172.16.0.0/24"
	AliCloudVSwitchDescription string = "This vswitch is created by alicloud-sdk-wrapper"

	// security group
	AliCloudSecurityGroupName        string = "sg-alicloud-sdk-wrapper"
	AliCloudSecurityGroupDescription string = "This securityGroup is created by alicloud-sdk-wrapper"

	// security group rule
	AliCloudSecurityGroupRuleIpProtocol   string = "tcp"
	AliCloudSecurityGroupRuleNicType      string = "intranet"
	AliCloudSecurityGroupRuleSourceCidrIp string = "0.0.0.0/0"
	AliCloudSecurityGroupRulePolicy       string = "accept"
	AliCloudSecurityGroupRulePriority     string = "1"
	AliCloudSecurityGroupRuleDescription  string = "This sg-rule is created by alicloud-sdk-wrapper"

	// instance
	AliCloudInstanceType                 string = "ecs.c5.large"
	AliCloudImageId                      string = "ubuntu_16_0402_64_20G_alibase_20180409.vhd"
	AliCloudSystemDiskCategory           string = "cloud_efficiency"
	AliCloudSystemDiskSize               string = "40"
	AliCloudInstanceName                 string = "i-alicloud-sdk-wrapper"
	AliCloudInstanceDescription          string = "This instance is created by alicloud-sdk-wrapper"
	AliCloudInternetChargeType           string = "PayByBandwidth"
	AliCloudInternetMaxBandwidthOut      int    = 1
	AliCloudInstanceHostName             string = "alicloud-sdk-wrapper"
	AliCloudInstancePassword             string = "Pass@w0rd"
	AliCloudInstanceChargeTypePrePaid    string = "PrePaid"
	AliCloudInstanceChargeTypePostPaid   string = "PostPaid"
	AliCloudInstanceChargeTypePeriodUnit string = "Month"
	AliCloudInstanceChargeTypePeriod     int    = 1
	AliCloudInstanceAmount               int    = 1
	AliCloudIoOptimized                  string = "optimized"
	AliCloudDryRun                       bool   = false
	AliCloudUniqueSuffix                 bool   = true
	AliCloudPageSize                     int    = 100

	// error messages
	AliCloudErrorMessageVPC         string = "no vpc"
	AliCloudErrorMessageInvalidData string = "invalid data"
)
