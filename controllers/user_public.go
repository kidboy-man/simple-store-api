package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"
)

// Operations about object
type UserPublicController struct {
	BaseController
	userUcase usecase.UserUsecase
}

func (c *UserPublicController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Register
// @Description register
// @Summary register
// @Success 200
// @Failure 403
// @Param params body datatransfers.RegisterRequest true "body of this request"
// @router /register [post]
func (c *UserPublicController) Register(params *datatransfers.RegisterRequest) *JSONResponse {
	err := c.userUcase.Register(params)
	return c.ReturnJSONResponse(nil, err)
}

// @Title Login
// @Description login
// @Summary login
// @Success 200
// @Failure 403
// @Param params body datatransfers.LoginRequest true "body of this request"
// @router /login [post]
func (c *UserPublicController) Login(params *datatransfers.LoginRequest) *JSONResponse {
	user, err := c.userUcase.Login(params)
	return c.ReturnJSONResponse(user, err)
}
