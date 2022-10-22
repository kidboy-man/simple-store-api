package repository

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VariantRepository interface {
	Create(variant *models.Variant, db *gorm.DB) (err error)
	Delete(variant *models.Variant, db *gorm.DB) (err error)
	GetAll(params *datatransfers.VariantQueryParams) (variants []*models.Variant, cnt int64, err error)
	GetByID(variantID int) (variant *models.Variant, err error)
	Update(variant *models.Variant, db *gorm.DB) (err error)
}
type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepository{db: db}
}

func (r *variantRepository) GetAll(params *datatransfers.VariantQueryParams) (variants []*models.Variant, cnt int64, err error) {
	qs := r.db
	if params.IsPublic {
		qs = qs.Where("is_active = ?", true)
	}

	if params.ProductID > 0 {
		qs = qs.Where("product_id = ?", params.ProductID)
	}

	if !params.IsWithoutCount {
		err = qs.Model(&models.Variant{}).Count(&cnt).Error
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

	err = qs.Find(&variants).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *variantRepository) GetByID(variantID int) (variant *models.Variant, err error) {
	qs := r.db.Where("id = ?", variantID)
	err = qs.Preload("Variants").First(&variant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusInternalServerError,
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

func (r *variantRepository) Create(variant *models.Variant, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(variant).Create(variant).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *variantRepository) Update(variant *models.Variant, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(variant).Updates(variant)
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

func (r *variantRepository) Delete(variant *models.Variant, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(variant).Delete(variant)
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
