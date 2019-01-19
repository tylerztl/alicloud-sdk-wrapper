package services

import (
	"zig-cloud/services/cloudprovider/alicloud"
	"zig-cloud/commons"
)

// define all kinds of services

type Provider interface {
	CreateInstance(request Request) Response
	RunInstances(request Request) Response
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