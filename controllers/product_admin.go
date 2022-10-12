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
type ProductAdminController struct {
	beego.Controller
	productUcase usecase.ProductUsecase
}

func (c *ProductAdminController) Prepare() {
	c.productUcase = usecase.NewProductUsecase(conf.AppConfig.DbClient)
}

// @Title Get All Products
// @Summary Get All Products
// @Description get all products
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *ProductAdminController) GetAll(limit, page int) (response JSONResponse) {
	product, totalData, err := c.productUcase.GetAll(&datatransfers.ProductQueryParams{
		BaseQueryParams: datatransfers.BaseQueryParams{
			Limit:    limit,
			Page:     page,
			Offset:   utils.CalculateOffset(limit, page),
			IsPublic: false,
		},
	})

	response.SetPagination(c.Ctx, totalData, limit, page)
	response.ReturnJSONResponse(product, err)
	return
}

// @Title Get Product
// @Summary Get Product
// @Description Get product
// @Success 200
// @Failure 403
// @Param productID path int true "id of the product to update"
// @router /:productID [get]
func (c *ProductAdminController) GetProduct(productID uint) (response JSONResponse) {
	product, err := c.productUcase.GetByID(productID)
	response.ReturnJSONResponse(product, err)
	return
}

// @Title Create Product
// @Summary Create Product
// @Description create product
// @Success 200
// @Failure 403
// @Param params body models.Product true "body of this request"
// @router / [post]
func (c *ProductAdminController) CreateProduct(params *models.Product) (response JSONResponse) {
	err := c.productUcase.Create(params)
	response.ReturnJSONResponse(params, err)
	return
}

// @Title Update Product
// @Summary Update Product
// @Description update product
// @Success 200
// @Failure 403
// @Param productID path int true "id of the product to update"
// @Param params body models.Product true "body of this request"
// @router /:productID [put]
func (c *ProductAdminController) UpdateProduct(productID uint, params *models.Product) (response JSONResponse) {
	// TODO: validate params
	params.ID = productID
	err := c.productUcase.Update(params)
	response.ReturnJSONResponse(params, err)
	return
}

// @Title Delete Product
// @Summary Delete Product
// @Description delete product
// @Success 200
// @Failure 403
// @Param productID path int true "id of the product to update"
// @router /:productID [delete]
func (c *ProductAdminController) DeleteProduct(productID uint) (response JSONResponse) {
	err := c.productUcase.Delete(&models.Product{ID: productID})
	response.ReturnJSONResponse(nil, err)
	return
}
