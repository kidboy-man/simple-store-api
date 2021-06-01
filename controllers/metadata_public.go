package controllers

import (
	"net/http"
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
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @Success 200
// @Failure 403
// @router / [get]
func (c *MetadataPublicController) GetAll(limit, page int) Response {
	// dummy errors
	errCust := &CustomError{
		Code:    111,
		Status:  http.StatusInternalServerError,
		Message: "something went wrong",
	}
	metadata, _, err := c.metadataUcase.GetAll(&datatransfers.ListQueryParams{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		errCust.Message = err.Error()
		return ReturnError(errCust)
	}
	return ReturnOK(metadata)
}