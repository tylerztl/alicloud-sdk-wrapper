package services

import (
	"zig-cloud/services/cloudprovider/alicloud"
	"zig-cloud/commons"
)

// define all kinds of services

type Provider interface {
	CreateInstance(request *commons.CreateInstanceRequest) (*commons.CreateInstanceResponse, error)
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