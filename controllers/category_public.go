package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type CategoryPublicController struct {
	beego.Controller
	categoryUcase usecase.CategoryUsecase
}

func (c *CategoryPublicController) Prepare() {
	c.categoryUcase = usecase.NewCategoryUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Categories
// @Summary Get All Categories
// @Summary Get All Categories
// @Description get only the active categories
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *CategoryPublicController) GetAll(limit, page int) JSONResponse {
	category, _, err := c.categoryUcase.GetAll(&datatransfers.ListQueryParams{
		Limit:    limit,
		Page:     page,
		IsPublic: true,
	})

	return ReturnJSONResponse(category, err)
}
