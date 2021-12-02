package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type UserPublicController struct {
	beego.Controller
	userUcase usecase.UserUsecase
}

func (c *UserPublicController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Register
// @Description register
// @Success 200
// @Failure 403
// @Param params body datatransfers.RegisterReques true "body of this request"
// @router /register [post]
func (c *UserPublicController) Register(params *datatransfers.RegisterRequest) JSONResponse {
	err := c.userUcase.Register(params)
	return ReturnJSONResponse(nil, err)
}
