package usecase

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	repository "simple-store-api/repositories"

	"gorm.io/gorm"
)

type ProductUsecase interface {
	Create(product *models.Product) (err error)
	Delete(product *models.Product) (err error)
	GetAll(param *datatransfers.ProductQueryParams) (products []*models.Product, cnt int64, err error)
	GetByID(productID uint) (product *models.Product, err error)
	Update(product *models.Product) (err error)
}

type productUsecase struct {
	db           *gorm.DB
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductUsecase(db *gorm.DB) ProductUsecase {
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	return &productUsecase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		db:           db,
	}
}

func (u *productUsecase) GetAll(param *datatransfers.ProductQueryParams) (products []*models.Product, cnt int64, err error) {
	products, cnt, err = u.productRepo.GetAll(param)
	return
}

func (u *productUsecase) GetByID(productID uint) (product *models.Product, err error) {
	product, err = u.productRepo.GetByID(productID)
	return
}

func (u *productUsecase) Create(product *models.Product) (err error) {
	if product.Category == nil {
		err = &datatransfers.CustomError{
			Code:   constants.BadRequesErrCode, // TODO: change Error Code
			Status: http.StatusBadRequest,
		}
	}

	category, err := u.categoryRepo.GetByID(product.Category.ID)
	if err != nil {
		return
	}

	product.Category = category
	err = u.productRepo.Create(product, u.db)
	return
}

func (u *productUsecase) Update(product *models.Product) (err error) {
	if product.Category != nil {
		category, errCategory := u.categoryRepo.GetByID(product.Category.ID)
		if errCategory != nil {
			return errCategory
		}

		product.Category = category
	}

	err = u.productRepo.Update(product, u.db)
	return
}

func (u *productUsecase) Delete(product *models.Product) (err error) {
	err = u.productRepo.Delete(product, u.db)
	return
}
