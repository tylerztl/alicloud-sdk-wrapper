package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/alicloud"
)


func TestProvider(t *testing.T) {
	var provider services.Provider
	provider = new(alicloud.CloudProvider)
	t.Log(provider)
}
