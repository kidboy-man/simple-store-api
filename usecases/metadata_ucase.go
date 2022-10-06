package usecase

import (
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	repository "simple-store-api/repositories"

	"gorm.io/gorm"
)

type MetadataUsecase interface {
	Create(metadata *models.Metadata) (err error)
	Delete(metadata *models.Metadata) (err error)
	GetAll(param *datatransfers.BaseQueryParams) (metadatas []*models.Metadata, cnt int64, err error)
	GetByID(metadataID int) (metadata *models.Metadata, err error)
	GetByType(slug string) (metadata *models.Metadata, err error)
	Update(metadata *models.Metadata) (err error)
}

type metadataUsecase struct {
	db           *gorm.DB
	metadataRepo repository.MetadataRepository
}

func NewMetadataUsecase(db *gorm.DB) MetadataUsecase {
	metadataRepo := repository.NewMetadataRepository(db)
	return &metadataUsecase{
		metadataRepo: metadataRepo,
		db:           db,
	}
}

func (u *metadataUsecase) GetAll(param *datatransfers.BaseQueryParams) (metadatas []*models.Metadata, cnt int64, err error) {
	metadatas, cnt, err = u.metadataRepo.GetAll(param)
	return
}

func (u *metadataUsecase) GetByID(metadataID int) (metadata *models.Metadata, err error) {
	metadata, err = u.metadataRepo.GetByID(metadataID)
	return
}

func (u *metadataUsecase) GetByType(metadataType string) (metadata *models.Metadata, err error) {
	metadata, err = u.metadataRepo.GetByType(metadataType)
	return
}

func (u *metadataUsecase) Create(metadata *models.Metadata) (err error) {
	err = u.metadataRepo.Create(metadata, u.db)
	return
}

func (u *metadataUsecase) Update(metadata *models.Metadata) (err error) {
	err = u.metadataRepo.Update(metadata, u.db)
	return
}

func (u *metadataUsecase) Delete(metadata *models.Metadata) (err error) {
	err = u.metadataRepo.Delete(metadata, u.db)
	return
}
