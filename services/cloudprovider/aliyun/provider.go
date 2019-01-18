package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/services"
	"zig-cloud/helpers"
)

type CloudProvider struct {

}

func (provider *CloudProvider) CreateInstance(request services.Request) services.Response {
	createInstanceRequest := &ecs.CreateInstanceRequest{}
	client := provider.GetClient()
	response, err := client.CreateInstance(createInstanceRequest)
	if err == nil {
		return helpers.GetServiceResponse(response)
	}else{
		return services.Response{}
	}
}

func (provider *CloudProvider) RunInstances(request services.Request) services.Response {
	client := provider.GetClient()
	runInstancesRequest := ecs.RunInstancesRequest{}
	client.RunInstances(&runInstancesRequest)
	return services.Response{}
}

func (provider *CloudProvider) GetClient()  *ecs.Client {
	regionId := ""
	accessKeyId := "76PKIccqlFZERf7X"
	accessKeySecret := "6QHPmxyEf80WqUNinteE1WWWqoz8IW"
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId,accessKeySecret)
	if err == nil {
		return client
	}else {
		return nil;
	}
}