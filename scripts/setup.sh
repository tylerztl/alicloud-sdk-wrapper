#!/usr/bin/env bash

# Download dependent packages

go get github.com/astaxie/beego
go get github.com/aliyun/alibaba-cloud-sdk-go
go get github.com/smartystreets/goconvey

# Download bee command tool

go get -u github.com/beego/bee

bee run -downdoc=true -gendoc true