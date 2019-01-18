package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/aliyun"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var regionId string = "cn-hangzhou"
var accessKeyId string = "76PKIccqlFZERf7X"
var accessKeySecret string = "6QHPmxyEf80WqUNinteE1WWWqoz8IW"


func TestAliyunClient(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	if err != nil {
		t.Error(err)
	}
	request := new(ecs.DescribeUserDataRequest)
	request.InstanceId = "i-rj99oa9orptlgqdhyftu"
	request.RegionId = "us-west-1"
	response, err := client.DescribeUserData(request)

	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}

func TestAliyunCreateInstance(t *testing.T) {
	var provider services.Provider
	provider = new(aliyun.CloudProvider)
	request := services.Request{}
	provider.RunInstances(request)
	t.Log("hello world")
}
