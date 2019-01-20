package services

import (
	"testing"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"zig-cloud/helpers"
	"zig-cloud/commons"
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

	authorizeSecurityGroupRequest := ecs.CreateAuthorizeSecurityGroupRequest();
	authorizeSecurityGroupRequest.SecurityGroupId = response.SecurityGroupId
	authorizeSecurityGroupRequest.NicType = "intranet"
	authorizeSecurityGroupRequest.IpProtocol = "tcp"
	authorizeSecurityGroupRequest.Policy = "accept"
	authorizeSecurityGroupRequest.PortRange = "80/80"
	authorizeSecurityGroupRequest.Priority = "1"
	authorizeSecurityGroupRequest.SourceCidrIp = "0.0.0.0/0"
	client.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
}


