package repository

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PriceRepository interface {
	Create(price *models.Price, db *gorm.DB) (err error)
	Delete(price *models.Price, db *gorm.DB) (err error)
	GetAll(params *datatransfers.PriceQueryParams) (prices []*models.Price, cnt int64, err error)
	GetByID(priceID uint) (price *models.Price, err error)
	Update(price *models.Price, db *gorm.DB) (err error)
}
type priceRepository struct {
	db *gorm.DB
}

func NewPriceRepository(db *gorm.DB) PriceRepository {
	return &priceRepository{db: db}
}

func (r *priceRepository) GetAll(params *datatransfers.PriceQueryParams) (prices []*models.Price, cnt int64, err error) {
	qs := r.db
	if params.IsPublic {
		qs = qs.Where("is_active = ?", true)
	}

	if params.VariantID > 0 {
		qs = qs.Where("product_id = ?", params.VariantID)
	}

	if !params.IsWithoutCount {
		err = qs.Model(&models.Price{}).Count(&cnt).Error
		if err != nil {
			err = &datatransfers.CustomError{
				Code:    constants.QueryInternalServerErrCode,
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			return
		}
	}

	if params.IsOnlyCount {
		return
	}

	if params.Limit > 0 {
		qs = qs.Limit(params.Limit)
	}

	if params.Offset > 0 {
		qs = qs.Offset(params.Offset)
	}

	err = qs.Find(&prices).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *priceRepository) GetByID(priceID uint) (price *models.Price, err error) {
	qs := r.db.Where("id = ?", priceID)
	err = qs.First(&price).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			return nil, err
		}

		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return nil, err
	}
	return
}

func (r *priceRepository) Create(price *models.Price, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(price).Create(price).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *priceRepository) Update(price *models.Price, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(price).Updates(price)
	err = row.Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if row.RowsAffected == 0 {
		err = &datatransfers.CustomError{
			Code:    constants.QueryNotFoundErrCode,
			Status:  http.StatusNotFound,
			Message: gorm.ErrRecordNotFound.Error(),
		}
	}
	return
}

func (r *priceRepository) Delete(price *models.Price, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(price).Delete(price)
	err = row.Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if row.RowsAffected == 0 {
		err = &datatransfers.CustomError{
			Code:    constants.QueryNotFoundErrCode,
			Status:  http.StatusNotFound,
			Message: gorm.ErrRecordNotFound.Error(),
		}
	}
	return
}
