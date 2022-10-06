package repository

import (
	"net/http"
	"simple-store-api/datatransfers"
	"simple-store-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MetadataRepository interface {
	Create(metadata *models.Metadata, db *gorm.DB) (err error)
	Delete(metadata *models.Metadata, db *gorm.DB) (err error)
	GetAll(params *datatransfers.BaseQueryParams) (metadatas []*models.Metadata, cnt int64, err error)
	GetByID(metadataID int) (metadata *models.Metadata, err error)
	GetByType(metadataType string) (metadata *models.Metadata, err error)
	Update(metadata *models.Metadata, db *gorm.DB) (err error)
}
type metadataRepository struct {
	db *gorm.DB
}

func NewMetadataRepository(db *gorm.DB) MetadataRepository {
	return &metadataRepository{db: db}
}

func (r *metadataRepository) GetAll(params *datatransfers.BaseQueryParams) (metadatas []*models.Metadata, cnt int64, err error) {
	qs := r.db.Where("deleted_at ISNULL")
	err = qs.Model(&models.Metadata{}).Count(&cnt).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    0,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	if params.Limit > 0 {
		qs = qs.Limit(params.Limit)
	}

	if params.Offset > 0 {
		qs = qs.Offset(params.Offset)
	}

	err = qs.Find(&metadatas).Error
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    0,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return
}

func (r *metadataRepository) GetByID(metadataID int) (metadata *models.Metadata, err error) {
	qs := r.db.Where("id = ?", metadataID)
	err = qs.First(metadata).Error
	return
}

func (r *metadataRepository) GetByType(metadataType string) (metadata *models.Metadata, err error) {
	qs := r.db.Where("type = ?", metadataType)
	err = qs.First(metadata).Error
	return
}

func (r *metadataRepository) Create(metadata *models.Metadata, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(metadata).Create(metadata).Error
	return
}

func (r *metadataRepository) Update(metadata *models.Metadata, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(metadata).Updates(metadata)
	if row.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	err = row.Error
	return
}

func (r *metadataRepository) Delete(metadata *models.Metadata, db *gorm.DB) (err error) {
	row := db.Omit(clause.Associations).Model(metadata).Delete(metadata)
	if row.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
		return
	}

	err = row.Error
	return
}
