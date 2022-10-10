package repository

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	Create(category *models.Category, db *gorm.DB) (err error)
	Delete(category *models.Category, db *gorm.DB) (err error)
	GetAll(params *datatransfers.CategoryQueryParams) (categories []*models.Category, cnt int64, err error)
	GetByID(categoryID uint) (category *models.Category, err error)
	Update(category *models.Category, db *gorm.DB) (err error)
}
type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(params *datatransfers.CategoryQueryParams) (categories []*models.Category, cnt int64, err error) {
	qs := r.db
	if params.IsPublic {
		qs = qs.Where("is_active = ?", true)
	}

	if !params.IsWithoutCount {
		err = qs.Model(&models.Category{}).Count(&cnt).Error
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

	err = qs.Find(&categories).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *categoryRepository) GetByID(categoryID uint) (category *models.Category, err error) {
	qs := r.db.Where("id = ?", categoryID)
	err = qs.First(category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &datatransfers.CustomError{
				Code:    constants.QueryNotFoundErrCode,
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			return
		}

		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *categoryRepository) Create(category *models.Category, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(category).Create(category).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.QueryInternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *categoryRepository) Update(category *models.Category, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(category).Updates(category)
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
			Status:  http.StatusInternalServerError,
			Message: gorm.ErrRecordNotFound.Error(),
		}
	}
	return
}

func (r *categoryRepository) Delete(category *models.Category, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(category).Delete(category)
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
			Status:  http.StatusInternalServerError,
			Message: gorm.ErrRecordNotFound.Error(),
		}
	}
	return
}
