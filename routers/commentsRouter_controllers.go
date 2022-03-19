package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryPublicController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryPublicController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("limit"),
				param.New("page"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:MetadataPublicController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:MetadataPublicController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("limit"),
				param.New("page"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:UserPublicController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:UserPublicController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:UserPublicController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:UserPublicController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
