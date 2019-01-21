package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/helpers"
	"zig-cloud/commons"
	"zig-cloud/services"
	"github.com/astaxie/beego"
)

type CloudProvider struct {
	config *services.AliCloudConfig
}

func (provider *CloudProvider) ConfigureClient(config *services.AliCloudConfig) {
	provider.config = config
}

func (provider *CloudProvider) CreateVpc(request *commons.CreateVpcRequest) (*commons.CreateVpcResponse, error) {
	return nil,nil
}

func (provider *CloudProvider) CreateVSwitch(request *commons.CreateVSwitchRequest) (*commons.CreateVpcResponse, error) {
	return nil,nil
}

func (provider *CloudProvider) CreateSecurityGroup(request *commons.CreateSecurityGroupRequest) (*commons.CreateSecurityGroupResponse, error) {
	return nil,nil
}

func (provider *CloudProvider) RunInstances(request *commons.RunInstancesRequest) (*commons.RunInstancesResponse, error) {
	client := provider.GetClient()
	runInstancesRequest := ecs.CreateRunInstancesRequest()
	helpers.TransferValuesBetweenRequest(runInstancesRequest,request)
	response, err := client.RunInstances(runInstancesRequest)
	if err == nil {
		return helpers.GetRunInstancesResponse(response), nil
	} else {
		return nil, err
	}
}

func (provider *CloudProvider) GetClient()  *ecs.Client {
	config := provider.config
	client, err := ecs.NewClientWithAccessKey(config.RegionId, config.AccessKeyId, config.AccessKeySecret)
	if err == nil {
		return client
	}else {
		beego.Error(err)
		return nil;
	}
}