package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	usecase "simple-store-api/usecases"
	"simple-store-api/utils"
)

// Operations about object
type PriceAdminController struct {
	BaseController
	priceUcase usecase.PriceUsecase
}

func (c *PriceAdminController) Prepare() {
	c.priceUcase = usecase.NewPriceUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Prices
// @Summary Get All Prices
// @Description get all prices
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *PriceAdminController) GetAll(limit, page int) *JSONResponse {
	price, totalData, err := c.priceUcase.GetAll(&datatransfers.PriceQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: false,
		},
	})

	return c.ReturnJSONListResponse(price, totalData, limit, page, err)
}

// @Title Get Price
// @Summary Get Price
// @Description Get price
// @Success 200
// @Failure 403
// @Param priceID path int true "id of the price to update"
// @router /:priceID [get]
func (c *PriceAdminController) GetPrice(priceID uint) *JSONResponse {
	price, err := c.priceUcase.GetByID(priceID)
	return c.ReturnJSONResponse(price, err)
}

// @Title Create Price
// @Summary Create Price
// @Description create price
// @Success 200
// @Failure 403
// @Param params body models.Price true "body of this request"
// @router / [post]
func (c *PriceAdminController) CreatePrice(params *models.Price) *JSONResponse {
	err := c.priceUcase.Create(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Update Price
// @Summary Update Price
// @Description update price
// @Success 200
// @Failure 403
// @Param priceID path int true "id of the price to update"
// @Param params body models.Price true "body of this request"
// @router /:priceID [put]
func (c *PriceAdminController) UpdatePrice(priceID uint, params *models.Price) *JSONResponse {
	// TODO: validate params
	params.ID = priceID
	err := c.priceUcase.Update(params)
	return c.ReturnJSONResponse(params, err)
}

// @Title Delete Price
// @Summary Delete Price
// @Description delete price
// @Success 200
// @Failure 403
// @Param priceID path int true "id of the price to update"
// @router /:priceID [delete]
func (c *PriceAdminController) DeletePrice(priceID uint) *JSONResponse {
	err := c.priceUcase.Delete(&models.Price{ID: priceID})
	return c.ReturnJSONResponse(nil, err)
}
