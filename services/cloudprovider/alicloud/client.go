package alicloud

import (
	"os"
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
