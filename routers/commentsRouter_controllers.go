package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "DeleteInstance",
            Router: `/instance/delete/:instanceId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "StopInstance",
            Router: `/instance/stop/:instanceId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "RunInstances",
            Router: `/instances`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "DescribeInstances",
            Router: `/instances`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "DescribeRegions",
            Router: `/regions`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "CreateSecurityGroup",
            Router: `/securitygroup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "AuthorizeSecurityGroup",
            Router: `/securitygroup/ingress-rule`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "CreateVPC",
            Router: `/vpc`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "CreateVSwitch",
            Router: `/vswitch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:CloudController"],
        beego.ControllerComments{
            Method: "DescribeZones",
            Router: `/zones`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zig-cloud/controllers:UserController"] = append(beego.GlobalControllerRouter["zig-cloud/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
