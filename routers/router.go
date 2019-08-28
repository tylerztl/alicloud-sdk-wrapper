// @APIVersion 1.0.0
// @Title Ali Cloud API
// @Description Manage IaaS Resources
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"alicloud-sdk-wrapper/commons"
	"alicloud-sdk-wrapper/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/alicloud",
			beego.NSInclude(
				&controllers.CloudController{Provider: commons.AliCloud},
			),
		),
	)
	beego.AddNamespace(ns)
}
