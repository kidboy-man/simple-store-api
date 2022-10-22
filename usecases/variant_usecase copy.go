package usecase

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	repository "simple-store-api/repositories"

	"gorm.io/gorm"
)

type PriceUsecase interface {
	Create(price *models.Price) (err error)
	Delete(price *models.Price) (err error)
	GetAll(param *datatransfers.PriceQueryParams) (prices []*models.Price, cnt int64, err error)
	GetByID(priceID uint) (price *models.Price, err error)
	Update(price *models.Price) (err error)
}

type priceUsecase struct {
	db          *gorm.DB
	priceRepo   repository.PriceRepository
	variantRepo repository.VariantRepository
}

func NewPriceUsecase(db *gorm.DB) PriceUsecase {
	priceRepo := repository.NewPriceRepository(db)
	variantRepo := repository.NewVariantRepository(db)
	return &priceUsecase{
		priceRepo:   priceRepo,
		variantRepo: variantRepo,
		db:          db,
	}
}

func (u *priceUsecase) GetAll(param *datatransfers.PriceQueryParams) (prices []*models.Price, cnt int64, err error) {
	prices, cnt, err = u.priceRepo.GetAll(param)
	return
}

func (u *priceUsecase) GetByID(priceID uint) (price *models.Price, err error) {
	price, err = u.priceRepo.GetByID(priceID)
	return
}

func (u *priceUsecase) Create(price *models.Price) (err error) {
	if price.Variant == nil {
		err = &datatransfers.CustomError{
			Code:   constants.BadRequestErrCode, // TODO: change Error Code
			Status: http.StatusBadRequest,
		}
		return
	}

	variant, err := u.variantRepo.GetByID(price.Variant.ID)
	if err != nil {
		return
	}

	price.Variant = variant
	err = u.priceRepo.Create(price, u.db)
	return
}

func (u *priceUsecase) Update(price *models.Price) (err error) {
	if price.Variant != nil {
		variant, errVariant := u.variantRepo.GetByID(price.Variant.ID)
		if errVariant != nil {
			return errVariant
		}

		price.Variant = variant
	}

	err = u.priceRepo.Update(price, u.db)
	return
}

func (u *priceUsecase) Delete(price *models.Price) (err error) {
	err = u.priceRepo.Delete(price, u.db)
	return
}
