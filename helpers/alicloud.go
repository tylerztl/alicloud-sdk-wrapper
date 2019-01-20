package helpers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/commons"
)

func GetCreateInstanceResponse(response *ecs.CreateInstanceResponse) *commons.CreateInstanceResponse {
	return &commons.CreateInstanceResponse{}
}

func GetRunInstancesResponse(response *ecs.RunInstancesResponse) *commons.RunInstancesResponse {
	return &commons.RunInstancesResponse{}
}