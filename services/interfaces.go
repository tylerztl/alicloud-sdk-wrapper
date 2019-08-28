package services

import (
	"alicloud-sdk-wrapper/commons"
	"alicloud-sdk-wrapper/services/cloudprovider/alicloud"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// define all kinds of services
type Provider interface {
	CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error)
	CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVSwitchResponse, error)
	CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error)
	AuthorizeSecurityGroup(request *commons.AuthorizeSecurityGroupRequest) (*commons.AuthorizeSecurityGroupResponse, error)
	RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error)
	StopInstance(instanceId string) (*commons.StopInstanceResponse, error)
	DeleteInstance(instanceId string) (*commons.DeleteInstanceResponse, error)
	DescribeInstances(request *commons.DescribeInstancesRequest) (*commons.DescribeInstancesResponse, error)
	DescribeRegions() (*ecs.DescribeRegionsResponse, error)
	DescribeZones(regionId string) ([]*commons.DescribeZonesResponse, error)
}

func GetProviderByType(providerType commons.CloudProvider) Provider {
	switch providerType {
	case commons.AliCloud:
		return alicloud.NewAliCloudProvider()
	default:
		return alicloud.NewAliCloudProvider()
	}
}
