package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/alicloud"
)


func TestAliCloudCreateInstance(t *testing.T) {
	var provider services.Provider
	provider = new(alicloud.CloudProvider)
	request := services.Request{}
	provider.RunInstances(request)
	t.Log("hello world")
}
