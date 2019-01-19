package services

import (
	"testing"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"zig-cloud/helpers"
	"zig-cloud/commons"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func TestCreateSecurityGroup(t *testing.T) {
	client := GetClient()
	request := ecs.CreateCreateSecurityGroupRequest()
	request.SecurityGroupName = helpers.GenerateSecurityGroupName()
	request.Description = commons.AliCloudSecurityGroupDescription
	response, err := client.CreateSecurityGroup(request)

	if err != nil {
		fmt.Println(">>> encounter request errors")
		t.Error(err)
	}
	fmt.Printf("The security group ID is %s", response.SecurityGroupId)
	t.Log(response.GetHttpContentString())
}


func TestCreateSecurityGroupRule(t *testing.T) {
	client := GetClient()
	securityGroupRule := ecs.CreateModifySecurityGroupRuleRequest()
	request :=  requests.NewCommonRequest()
	request.Product = securityGroupRule.GetProduct()
	response, err := client.ProcessCommonRequest(request)

	if err == nil {
		fmt.Println(response.GetHttpContentString())
	}
}