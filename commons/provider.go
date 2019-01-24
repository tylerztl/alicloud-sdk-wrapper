package commons

type CloudProvider string

const (
	AliCloud = CloudProvider("alicloud")
)

const (
	RegionId        string = "regionid"
	AccessKeyId     string = "accesskeyid"
	AccessKeySecret string = "accesskeysecret"
	UseEnv          string = "useenv"
)
