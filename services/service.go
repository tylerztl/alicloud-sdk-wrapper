package services

import (
	"zig-cloud/services/cloudprovider/alicloud"
	"zig-cloud/commons"
	"github.com/astaxie/beego"
	"os"
)

// define all kinds of services

type Provider interface {
	CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error)
	CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error)
	CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error)
	RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse,error)
}

func NewAliCloudProvider() Provider {
	var RegionId,  AccessKeyId, AccessKeySecret string
	useEnv,err := beego.AppConfig.Bool(commons.UseEnv)

	if err == nil && useEnv {
		RegionId = os.Getenv(commons.RegionId)
		AccessKeyId = os.Getenv(commons.AccessKeyId)
		AccessKeySecret = os.Getenv(commons.AccessKeySecret)
	}else {
		RegionId = beego.AppConfig.String(commons.RegionId)
		AccessKeyId = beego.AppConfig.String(commons.AccessKeyId)
		AccessKeySecret = beego.AppConfig.String(commons.AccessKeySecret)
	}
	config := &AliCloudConfig{RegionId:RegionId,AccessKeyId:AccessKeyId,AccessKeySecret:AccessKeySecret}
	provider := &alicloud.CloudProvider{Config:config}
	return provider
}

func GetProviderByType(providerType string) Provider {
	switch providerType {
	case commons.CloudProviderAliCloud:
		return NewAliCloudProvider()
	default:
		return NewAliCloudProvider()
	}
}