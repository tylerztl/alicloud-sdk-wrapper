package services

import (
	"testing"
	"zig-cloud/services"
	"zig-cloud/services/cloudprovider/aliyun"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
)


func GetClient() *ecs.Client {
	//var regionId string = "cn-hangzhou"
	var regionId string = "us-west-1"
	var accessKeyId string = "76PKIccqlFZERf7X"
	var accessKeySecret string = "6QHPmxyEf80WqUNinteE1WWWqoz8IW"

	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	if err == nil {
		return client
	} else {
		fmt.Println(">>> encounter errors")
		return nil
	}
}

func TestHelloWorld(t *testing.T) {
	t.Log("Hello World")
}

func TestAliyunClient(t *testing.T) {

	client := GetClient()

	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = "i-rj99oa9orptlgqdhyftu"
	response, err := client.DescribeInstanceAttribute(request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.GetHttpContentString())
}

func TestAliyunCreateInstance(t *testing.T) {
	var provider services.Provider
	provider = new(aliyun.CloudProvider)
	request := services.Request{}
	provider.RunInstances(request)
	t.Log("hello world")
}
