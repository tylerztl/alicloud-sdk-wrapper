package services

import (
	"zig-cloud/commons"
	"zig-cloud/services/cloudprovider/alicloud"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// define all kinds of services
type Provider interface {
	CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error)
	CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error)
	CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error)
	AuthorizeSecurityGroup(request *commons.AuthorizeSecurityGroupRequest) (*commons.AuthorizeSecurityGroupResponse, error)
	RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error)
	DeleteInstance(instanceId string) (*commons.DeleteInstanceResponse, error)
	DescribeRegions() (*ecs.DescribeRegionsResponse, error)
	DescribeZones(regionId string) ([]*commons.DescribeZonesResponse, error)
}

func GetProviderByType(providerType commons.CloudProvider) Provider {
	switch providerType {
	case commons.AliCloud:
		return NewAliCloudProvider()
	default:
		return NewAliCloudProvider()
	}
}

func NewAliCloudProvider() Provider {
	return &alicloud.CloudProvider{Client: alicloud.GetClient()}
}
