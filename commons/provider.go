package commons

type CloudProvider string

const (
	AliCloud = CloudProvider("alicloud")
)

type AppConf string

const (
	RegionId        = AppConf("regionid")
	AccessKeyId     = AppConf("accesskeyid")
	AccessKeySecret = AppConf("accesskeysecret")
	SnapshotImageId = AppConf("snapshotimageid")
	UseEnv          = AppConf("useenv")
)

type CloudEnv string

const (
	EnvRegionId        = CloudEnv("ALICLOUD_REGION_ID")
	EnvAccessKeyId     = CloudEnv("ALICLOUD_ACCESS_KEY_ID")
	EnvAccessKeySecret = CloudEnv("ALICLOUD_ACCESS_KEY_SECRET")
	EnvSnapshotImageId = CloudEnv("ALICLOUD_SNAPSHOT_IMAGE_ID")
)
