package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/aliyun"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
)

func TestHelloWorld(t *testing.T) {
	t.Log("Hello World")
}

func TestAliyunClient(t *testing.T) {

	//var regionId string = "cn-hangzhou"
	var regionId string = "us-west-1"
	var accessKeyId string = "76PKIccqlFZERf7X"
	var accessKeySecret string = "6QHPmxyEf80WqUNinteE1WWWqoz8IW"

	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	if err != nil {
		fmt.Println(">>> encounter errors")
		t.Error(err)
	}

	/*
	request := new(ecs.DescribeUserDataRequest)
	request.InstanceId = "i-rj99oa9orptlgqdhyftu"
	request.RegionId = "us-west-1"
	*/
	request := ecs.DescribeUserDataRequest{}
	request.InstanceId = "i-rj99oa9orptlgqdhyftu"
	response, err := client.DescribeUserData(&request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.InstanceId)
}

func TestAliyunCreateInstance(t *testing.T) {
	var provider services.Provider
	provider = new(aliyun.CloudProvider)
	request := services.Request{}
	provider.RunInstances(request)
	t.Log("hello world")
}
