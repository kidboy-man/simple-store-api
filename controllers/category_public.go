package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"
)

// Operations about object
type CategoryPublicController struct {
	BaseController
	categoryUcase usecase.CategoryUsecase
}

func (c *CategoryPublicController) Prepare() {
	c.categoryUcase = usecase.NewCategoryUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Categories
// @Summary Get All Categories
// @Description get only the active categories
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *CategoryPublicController) GetAll(limit, page int) *JSONResponse {
	category, totalData, err := c.categoryUcase.GetAll(&datatransfers.CategoryQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: true,
		},
	})

	return c.ReturnJSONListResponse(category, totalData, limit, page, err)
}
