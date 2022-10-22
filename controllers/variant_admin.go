package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"
)

// Operations about object
type VariantAdminController struct {
	BaseController
	variantUcase usecase.VariantUsecase
}

func (c *VariantAdminController) Prepare() {
	c.variantUcase = usecase.NewVariantUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Variants
// @Summary Get All Variants
// @Description get all variants
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *VariantAdminController) GetAll(limit, page int) *JSONResponse {
	variant, totalData, err := c.variantUcase.GetAll(&datatransfers.VariantQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: false,
		},
	})

	return c.ReturnJSONListResponse(variant, totalData, limit, page, err)
}

// @Title Get Variant
// @Summary Get Variant
// @Description Get variant
// @Success 200
// @Failure 403
// @Param variantID path int true "id of the variant to update"
// @router /:variantID [get]
func (c *VariantAdminController) GetVariant(variantID uint) *JSONResponse {
	variant, err := c.variantUcase.GetByID(variantID)
	return c.ReturnJSONResponse(variant, err)
}

// @Title Create Variant
// @Summary Create Variant
// @Description create variant
// @Success 200
// @Failure 403
// @Param params body models.Variant true "body of this request"
// @router / [post]
func (c *VariantAdminController) CreateVariant(params *models.Variant) *JSONResponse {
	err := c.variantUcase.Create(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Update Variant
// @Summary Update Variant
// @Description update variant
// @Success 200
// @Failure 403
// @Param variantID path int true "id of the variant to update"
// @Param params body models.Variant true "body of this request"
// @router /:variantID [put]
func (c *VariantAdminController) UpdateVariant(variantID uint, params *models.Variant) *JSONResponse {
	// TODO: validate params
	params.ID = variantID
	err := c.variantUcase.Update(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Delete Variant
// @Summary Delete Variant
// @Description delete variant
// @Success 200
// @Failure 403
// @Param variantID path int true "id of the variant to update"
// @router /:variantID [delete]
func (c *VariantAdminController) DeleteVariant(variantID uint) *JSONResponse {
	err := c.variantUcase.Delete(&models.Variant{ID: variantID})
	return c.ReturnJSONResponse(nil, err)
}
