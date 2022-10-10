package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type CategoryAdminController struct {
	beego.Controller
	categoryUcase usecase.CategoryUsecase
}

func (c *CategoryAdminController) Prepare() {
	c.categoryUcase = usecase.NewCategoryUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Categories
// @Summary Get All Categories
// @Description get all categories
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *CategoryAdminController) GetAll(limit, page int) (response JSONResponse) {
	category, totalData, err := c.categoryUcase.GetAll(&datatransfers.CategoryQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: false,
		},
	})

	response.SetPagination(c.Ctx, totalData, limit, page)
	response.ReturnJSONResponse(category, err)
	return
}

// @Title Create Category
// @Summary Create Category
// @Description create category
// @Success 200
// @Failure 403
// @Param params body models.Category true "body of this request"
// @router / [post]
func (c *CategoryAdminController) CreateCategory(params *models.Category) (response JSONResponse) {
	// TODO: validate params
	err := c.categoryUcase.Create(params)
	response.ReturnJSONResponse(params, err)
	return
}
