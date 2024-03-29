# alicloud-sdk-wrapper

## Prerequisites

* Go 1.10+ installation or later
* **GOPATH** environment variable is set correctly
* Beego version 1.10.0 or later
* Govendor version 1.0.9 or later

## Getting Started
```
cd alicloud-sdk-wrapper
govendor sync
bee run -downdoc=true -gendoc true
```
* [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

## AliCloud

The providers to be used:

* [alicloud-ecs](https://github.com/terraform-alicloud-modules/terraform-alicloud-ecs-instance)
* [alicloud-vpc](https://github.com/terraform-alicloud-modules/terraform-alicloud-vpc)
* [alicloud-api](https://help.aliyun.com/document_detail/25485.html?spm=a2c4g.11186623.6.921.d74734b9JCtJ82#h2-url-15)

Related source codes:

- [terraform-provider-alicloud](https://github.com/terraform-providers/terraform-provider-alicloud)
- [alibaba-cloud-sdk-go](https://github.com/aliyun/alibaba-cloud-sdk-go)
- [openapi-core-nodejs-sdk](https://github.com/aliyun/openapi-core-nodejs-sdk)
