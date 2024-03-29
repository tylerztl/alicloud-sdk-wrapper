package services

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func GetClient() *ecs.Client {

	var regionId string = "cn-huhehaote"
	var accessKeyId string = "LTAIP5a8969b3WUho"
	var accessKeySecret string = "R94ZD14r8kfZlEZuk69mclpppxKSoCc"

	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	if err == nil {
		return client
	} else {
		fmt.Println(">>> encounter errors")
		return nil
	}
}

func TestAliCloudClient(t *testing.T) {

	client := GetClient()

	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = "i-hp38tm132y3ctkm9q65a"
	response, err := client.DescribeInstanceAttribute(request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.GetHttpContentString())
}

func TestCreateInstance(t *testing.T) {
	client := GetClient()
	request := ecs.CreateCreateInstanceRequest()
	request.ImageId = "ubuntu_16_0402_64_20G_alibase_20180409.vhd"
	request.InstanceName = "MyInstance"
	//request.SecurityGroupId = "sg-rj9f8xghzfk3jui835w8"
	request.InstanceType = "ecs.c5.large"
	//request.InstanceType = "ecs.t1.small"
	request.ClientToken = utils.GetUUIDV4()
	response, err := client.CreateInstance(request)
	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.GetHttpContentString())
}

func TestDeleteInstance(t *testing.T) {
	client := GetClient()
	request := ecs.CreateDeleteInstanceRequest()
	request.InstanceId = "i-hp38tm132y3ctkm9q65a"
	response, err := client.DeleteInstance(request)
	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.GetHttpContentString())
}

func TestDescribeRegions(t *testing.T) {
	client := GetClient()

	request := ecs.CreateDescribeRegionsRequest()

	response, err := client.DescribeRegions(request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	t.Log(response.GetHttpContentString())
}
