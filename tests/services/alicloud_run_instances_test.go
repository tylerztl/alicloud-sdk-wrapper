package services

import (
	"testing"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"zig-cloud/commons"
)

func TestAliCloudRunInstances(t *testing.T) {
	client := GetClient()

	runInstanceRequest := ecs.CreateRunInstancesRequest()
	runInstanceRequest.InstanceType = commons.AliCloudInstanceType

	response, err := client.RunInstances(runInstanceRequest)

	if err == nil {
		fmt.Println(response.GetHttpContentString())
	}else {
		t.Error(err)
	}
}
