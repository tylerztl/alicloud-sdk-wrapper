package services

import (
	"zig-cloud/services/cloudprovider/alicloud"
	"zig-cloud/commons"
)

// define all kinds of services

type Provider interface {
	CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error)
	CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVpcResponse, error)
	CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error)
	RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse,error)
}

func NewAliCloudProvider() Provider {
	provider := &alicloud.CloudProvider{}
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