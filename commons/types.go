package commons

type CreateVpcRequest struct {
	VpcName   string
	RegionId  string
	CidrBlock string
}

type CreateVpcResponse struct {
	VpcId string
}

type DescribeVpcAttributeRequest struct {
	RegionId string
	VpcId    string
}

type DescribeVpcAttributeResponse struct {
	VpcId        string
	RegionId     string
	Status       string
	VpcName      string
	CreationTime string
	CidrBlock    string
	VRouterId    string
}

type CreateVSwitchRequest struct {
	VSwitchName string
	VpcId       string
	ZoneId      string
	RegionId    string
	CidrBlock   string
}

type DescribeVSwitchAttributesRequest struct {
	RegionId  string
	VSwitchId string
}

type DescribeVSwitchAttributesResponse struct {
	VSwitchId               string
	VpcId                   string
	Status                  string
	CidrBlock               string
	ZoneId                  string
	AvailableIpAddressCount int
	VSwitchName             string
	CreationTime            string
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
	IpProtocol      string
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
	PeriodUnit         string
	Period             string
	Amount             int
	HostName           string
	Password           string
	TagKey             string
	TagValue           string
}

type Instance struct {
	InstanceId       string
	InstanceName     string
	HostName         string
	Status           string
	InstanceType     string
	CreationTime     string
	SecurityGroupIds []string
	PublicIpAddress  []string
}

type RunInstancesResponse struct {
	Instances []Instance
}

type DescribeInstancesRequest struct {
	RegionId    string
	InstanceIds string
}

type DescribeInstancesResponse struct {
	Instances []Instance
}

type StopInstanceResponse struct {
	RequestId string
}

type DeleteInstanceResponse struct {
	RequestId string
}

type DescribeZonesResponse struct {
	ZoneId    string
	LocalName string
}
