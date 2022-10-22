package usecase

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	repository "simple-store-api/repositories"

	"gorm.io/gorm"
)

type VariantUsecase interface {
	Create(variant *models.Variant) (err error)
	Delete(variant *models.Variant) (err error)
	GetAll(param *datatransfers.VariantQueryParams) (variants []*models.Variant, cnt int64, err error)
	GetByID(variantID uint) (variant *models.Variant, err error)
	Update(variant *models.Variant) (err error)
}

type variantUsecase struct {
	db          *gorm.DB
	variantRepo repository.VariantRepository
	productRepo repository.ProductRepository
}

func NewVariantUsecase(db *gorm.DB) VariantUsecase {
	variantRepo := repository.NewVariantRepository(db)
	productRepo := repository.NewProductRepository(db)
	return &variantUsecase{
		variantRepo: variantRepo,
		productRepo: productRepo,
		db:          db,
	}
}

func (u *variantUsecase) GetAll(param *datatransfers.VariantQueryParams) (variants []*models.Variant, cnt int64, err error) {
	variants, cnt, err = u.variantRepo.GetAll(param)
	return
}

func (u *variantUsecase) GetByID(variantID uint) (variant *models.Variant, err error) {
	variant, err = u.variantRepo.GetByID(variantID)
	return
}

func (u *variantUsecase) Create(variant *models.Variant) (err error) {
	if variant.Product == nil {
		err = &datatransfers.CustomError{
			Code:   constants.BadRequestErrCode, // TODO: change Error Code
			Status: http.StatusBadRequest,
		}
		return
	}

	product, err := u.productRepo.GetByID(variant.Product.ID)
	if err != nil {
		return
	}

	variant.Product = product
	err = u.variantRepo.Create(variant, u.db)
	return
}

func (u *variantUsecase) Update(variant *models.Variant) (err error) {
	if variant.Product != nil {
		product, errProduct := u.productRepo.GetByID(variant.Product.ID)
		if errProduct != nil {
			return errProduct
		}

		variant.Product = product
	}

	err = u.variantRepo.Update(variant, u.db)
	return
}

func (u *variantUsecase) Delete(variant *models.Variant) (err error) {
	err = u.variantRepo.Delete(variant, u.db)
	return
}
