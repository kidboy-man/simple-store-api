package controllers

import (
	"simple-store-api/conf"
	"simple-store-api/datatransfers"
	usecase "simple-store-api/usecases"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type MetadataPublicController struct {
	beego.Controller
	metadataUcase usecase.MetadataUsecase
}

func (c *MetadataPublicController) Prepare() {
	c.metadataUcase = usecase.NewMetadataUsecase(conf.AppConfig.DbClient)
}

// @Title GetAll
// @Description get all metadata
// @Summary get all metadata
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *MetadataPublicController) GetAll(limit, page int) (response JSONResponse) {
	metadata, _, err := c.metadataUcase.GetAll(&datatransfers.BaseQueryParams{
		Limit: limit,
		Page:  page,
	})

	response.ReturnJSONResponse(metadata, err)
	return
}
