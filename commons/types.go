package commons

type CreateVpcRequest struct {
	VpcName   string
	RegionId  string
	CidrBlock string
}

type CreateVpcResponse struct {
	VpcId string
}

type CreateVSwitchRequest struct {
	VSwitchName string
	VpcId       string
	ZoneId      string
	RegionId    string
	CidrBlock   string
}

type CreateVSwitchResponse struct {
	VSwitchId string `json:"VSwitchId" xml:"VSwitchId"`
}

type CreateSecurityGroupRequest struct {
	RegionId          string
	SecurityGroupName string
	VpcId             string
}

type CreateSecurityGroupResponse struct {
	SecurityGroupId string `json:"SecurityGroupId" xml:"SecurityGroupId"`
}

type AuthorizeSecurityGroupRequest struct {
	RegionId        string
	SecurityGroupId string
	PortRange       string
}

type AuthorizeSecurityGroupResponse struct {
	RequestId string
}

type RunInstancesRequest struct {
	RegionId           string
	ZoneId             string
	ImageId            string
	InstanceType       string
	SecurityGroupId    string
	VSwitchId          string
	InstanceChargeType string
	Amount             string
	Password           string
}

type RunInstancesResponse struct {
	Instances map[string]map[string]string
}

type DeleteInstanceResponse struct {
	RequestId string
}

type DescribeZonesResponse struct {
	ZoneId    string
	LocalName string
}
