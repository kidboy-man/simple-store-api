package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"],
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

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"],
        beego.ControllerComments{
            Method: "CreateCategory",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"],
        beego.ControllerComments{
            Method: "GetCategory",
            Router: "/:categoryID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("categoryID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"],
        beego.ControllerComments{
            Method: "UpdateCategory",
            Router: "/:categoryID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("categoryID", param.IsRequired, param.InPath),
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:CategoryAdminController"],
        beego.ControllerComments{
            Method: "DeleteCategory",
            Router: "/:categoryID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("categoryID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

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

    beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"],
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

    beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"],
        beego.ControllerComments{
            Method: "CreatePrice",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"],
        beego.ControllerComments{
            Method: "GetPrice",
            Router: "/:priceID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("priceID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"],
        beego.ControllerComments{
            Method: "UpdatePrice",
            Router: "/:priceID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("priceID", param.IsRequired, param.InPath),
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:PriceAdminController"],
        beego.ControllerComments{
            Method: "DeletePrice",
            Router: "/:priceID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("priceID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"],
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

    beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"],
        beego.ControllerComments{
            Method: "CreateProduct",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: "/:productID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("productID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"],
        beego.ControllerComments{
            Method: "UpdateProduct",
            Router: "/:productID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("productID", param.IsRequired, param.InPath),
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:ProductAdminController"],
        beego.ControllerComments{
            Method: "DeleteProduct",
            Router: "/:productID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("productID", param.IsRequired, param.InPath),
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

    beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"],
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

    beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"],
        beego.ControllerComments{
            Method: "CreateVariant",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"],
        beego.ControllerComments{
            Method: "GetVariant",
            Router: "/:variantID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("variantID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"],
        beego.ControllerComments{
            Method: "UpdateVariant",
            Router: "/:variantID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("variantID", param.IsRequired, param.InPath),
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"] = append(beego.GlobalControllerRouter["simple-store-api/controllers:VariantAdminController"],
        beego.ControllerComments{
            Method: "DeleteVariant",
            Router: "/:variantID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("variantID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

}
