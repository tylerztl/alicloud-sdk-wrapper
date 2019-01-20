package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/helpers"
	"zig-cloud/commons"
)

type CloudProvider struct {
	regionId string
	accessKeyId string
	accessKeySecret string
}

func (provider *CloudProvider) ConfigureClient() {

}

func (provider *CloudProvider) CreateInstance(request *commons.CreateInstanceRequest) (*commons.CreateInstanceResponse,error) {
	client := provider.GetClient()
	createInstanceRequest := ecs.CreateCreateInstanceRequest()
	response, err := client.CreateInstance(createInstanceRequest)
	if err == nil {
		return helpers.GetCreateInstanceResponse(response), nil
	}else{
		return nil, err
	}
}

func (provider *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	client := provider.GetClient()
	runInstancesRequest := ecs.CreateRunInstancesRequest()
	response, err := client.RunInstances(runInstancesRequest)
	if err == nil {
		return helpers.GetRunInstancesResponse(response), nil
	} else {
		return nil, err
	}
}

func (provider *CloudProvider) GetClient()  *ecs.Client {
	client, err := ecs.NewClientWithAccessKey(provider.regionId, provider.accessKeyId, provider.accessKeySecret)
	if err == nil {
		return client
	}else {
		return nil;
	}
}