package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"
)

// Operations about object
type CategoryAdminController struct {
	BaseController
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
func (c *CategoryAdminController) GetAll(limit, page int) *JSONResponse {
	category, totalData, err := c.categoryUcase.GetAll(&datatransfers.CategoryQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: false,
		},
	})

	return c.ReturnJSONListResponse(category, totalData, limit, page, err)
}

// @Title Get Category
// @Summary Get Category
// @Description Get category
// @Success 200
// @Failure 403
// @Param categoryID path int true "id of the category to update"
// @router /:categoryID [get]
func (c *CategoryAdminController) GetCategory(categoryID uint) *JSONResponse {
	category, err := c.categoryUcase.GetByID(categoryID)
	return c.ReturnJSONResponse(category, err)
}

// @Title Create Category
// @Summary Create Category
// @Description create category
// @Success 200
// @Failure 403
// @Param params body models.Category true "body of this request"
// @router / [post]
func (c *CategoryAdminController) CreateCategory(params *models.Category) *JSONResponse {
	err := c.categoryUcase.Create(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Update Category
// @Summary Update Category
// @Description update category
// @Success 200
// @Failure 403
// @Param categoryID path int true "id of the category to update"
// @Param params body models.Category true "body of this request"
// @router /:categoryID [put]
func (c *CategoryAdminController) UpdateCategory(categoryID uint, params *models.Category) *JSONResponse {
	// TODO: validate params
	params.ID = categoryID
	err := c.categoryUcase.Update(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Delete Category
// @Summary Delete Category
// @Description delete category
// @Success 200
// @Failure 403
// @Param categoryID path int true "id of the category to update"
// @router /:categoryID [delete]
func (c *CategoryAdminController) DeleteCategory(categoryID uint) *JSONResponse {
	err := c.categoryUcase.Delete(&models.Category{ID: categoryID})
	return c.ReturnJSONResponse(nil, err)
}
