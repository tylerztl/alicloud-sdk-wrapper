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
	VpcId       string
	VSwitchName string
	CidrBlock   string
	ZoneId      string
}

type CreateVSwitchResponse struct {
	VSwitchId string `json:"VSwitchId" xml:"VSwitchId"`
}

type CreateSecurityGroupRequest struct {
	SecurityGroupName string
	VpcId             string
}

type CreateSecurityGroupResponse struct {
	SecurityGroupId string `json:"SecurityGroupId" xml:"SecurityGroupId"`
}

type RunInstancesRequest struct {
	ImageId            string
	InstanceType       string
	InstanceChargeType string
	VpcId              string
	SecurityGroupId    string
	VSwitchId          string
	Amount             string
}

type RunInstancesResponse struct {
	Instances map[string]map[string]string
}

type DescribeZonesResponse struct {
	ZoneId    string
	LocalName string
}
