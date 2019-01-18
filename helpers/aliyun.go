package helpers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"zig-cloud/services"
)

func GetServiceResponse(response *ecs.CreateInstanceResponse) services.Response {
	return services.Response{}
}