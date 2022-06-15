package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"

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
func (c *CategoryPublicController) GetAll(limit, page int) (response JSONResponse) {
	category, totalData, err := c.categoryUcase.GetAll(&datatransfers.ListQueryParams{
		Limit:    limit,
		Page:     page,
		Offset:   utils.CalculateOffset(limit, page),
		IsPublic: true,
	})

	response.SetPagination(c.Ctx, totalData, limit, page)
	response.ReturnJSONResponse(category, err)
	return
}
