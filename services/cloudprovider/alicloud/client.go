package alicloud

import (
	"fmt"
	"os"
	"sync"
	"zig-cloud/commons"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/astaxie/beego"
)

type Config struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
}

type AliCloudClient struct {
	clientMutex     *sync.RWMutex
	config          *Config
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	ecsconn         *ecs.Client
	vpcconn         *vpc.Client
}

var client *AliCloudClient

func (c *Config) Client() *AliCloudClient {
	return &AliCloudClient{
		clientMutex:     new(sync.RWMutex),
		config:          c,
		RegionId:        c.RegionId,
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
	}
}

func init() {
	var RegionId, AccessKeyId, AccessKeySecret string
	useEnv, err := beego.AppConfig.Bool(commons.UseEnv)
	if err == nil && useEnv {
		RegionId = os.Getenv(commons.RegionId)
		AccessKeyId = os.Getenv(commons.AccessKeyId)
		AccessKeySecret = os.Getenv(commons.AccessKeySecret)
	} else {
		RegionId = beego.AppConfig.String(commons.RegionId)
		AccessKeyId = beego.AppConfig.String(commons.AccessKeyId)
		AccessKeySecret = beego.AppConfig.String(commons.AccessKeySecret)
	}

	config := Config{
		RegionId:        RegionId,
		AccessKeyId:     AccessKeyId,
		AccessKeySecret: AccessKeySecret,
	}
	client = config.Client()
}

func GetClient() *AliCloudClient {
	return client
}

func (a *AliCloudClient) WithEcsClient(do func(*ecs.Client) (interface{}, error)) (interface{}, error) {
	a.clientMutex.Lock()
	defer a.clientMutex.Unlock()

	if a.ecsconn == nil {
		ecsconn, err := ecs.NewClientWithAccessKey(a.RegionId, a.AccessKeyId, a.AccessKeySecret)
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the ECS client: %#v", err)
		}
		if _, err := ecsconn.DescribeRegions(ecs.CreateDescribeRegionsRequest()); err != nil {
			return nil, err
		}
		a.ecsconn = ecsconn
	}

	return do(a.ecsconn)
}

func (a *AliCloudClient) WithVpcClient(do func(*vpc.Client) (interface{}, error)) (interface{}, error) {
	a.clientMutex.Lock()
	defer a.clientMutex.Unlock()

	if a.vpcconn == nil {
		vpcconn, err := vpc.NewClientWithAccessKey(a.RegionId, a.AccessKeyId, a.AccessKeySecret)
		if err != nil {
			return nil, fmt.Errorf("unable to initialize the VPC client: %#v", err)
		}
		a.vpcconn = vpcconn
	}

	return do(a.vpcconn)
}
