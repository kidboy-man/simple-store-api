package repository

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product *models.Product, db *gorm.DB) (err error)
	Delete(product *models.Product, db *gorm.DB) (err error)
	GetAll(params *datatransfers.ProductQueryParams) (products []*models.Product, cnt int64, err error)
	GetByID(productID uint) (product *models.Product, err error)
	Update(product *models.Product, db *gorm.DB) (err error)
}
type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(params *datatransfers.ProductQueryParams) (products []*models.Product, cnt int64, err error) {
	qs := r.db
	if params.IsPublic {
		qs = qs.Where("is_active = ?", true)
	}

	if params.CategoryID > 0 {
		qs = qs.Where("category_id = ?", params.CategoryID)
	}

	if !params.IsWithoutCount {
		err = qs.Model(&models.Product{}).Count(&cnt).Error
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

	err = qs.Preload("Variants").Find(&products).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *productRepository) GetByID(productID uint) (product *models.Product, err error) {
	qs := r.db.Where("id = ?", productID)
	err = qs.Preload("Variants").First(&product).Error
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

func (r *productRepository) Create(product *models.Product, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(product).Create(product).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *productRepository) Update(product *models.Product, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(product).Updates(product)
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

func (r *productRepository) Delete(product *models.Product, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(product).Delete(product)
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
