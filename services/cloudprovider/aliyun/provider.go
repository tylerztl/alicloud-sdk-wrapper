package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type CloudProvider struct {

}

func (provider *CloudProvider) CreateInstance() interface{} {
	createInstanceRequest := &ecs.CreateInstanceRequest{}
	client := provider.GetClient()
	response, err := client.CreateInstance(createInstanceRequest)
	if err == nil {
		return response
	}else{
		return nil
	}
}

func (provider *CloudProvider) RunInstance() {
	client := provider.GetClient()
	runInstancesRequest := ecs.RunInstancesRequest{}
	client.RunInstances(&runInstancesRequest)
}

func (provider *CloudProvider) GetClient()  *ecs.Client {
	regionId := ""
	accessKeyId := ""
	accessKeySecret := ""
	client, err := ecs.NewClientWithAccessKey(regionId, accessKeyId,accessKeySecret)
	if err == nil {
		return client
	}else {
		return nil;
	}
}