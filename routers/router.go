// @APIVersion 1.0.0
// @Title Zig Cloud API
// @Description Manage IaaS Resources
// @Contact fanjiahe@zhigui.com
// @TermsOfServiceUrl http://zhigui.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"zig-cloud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/cloud",
			beego.NSInclude(
				&controllers.CloudController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
