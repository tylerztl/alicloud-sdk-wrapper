package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/aliyun"
)


func TestAliyunCreateInstance(t *testing.T){
	var provider services.Provider
	provider = new(aliyun.CloudProvider)
	request := services.Request{}
	provider.RunInstances(request)
	t.Log("hello world")
}
