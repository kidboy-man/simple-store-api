package usecase

import (
	"simple-store-api/datatransfers"
	"simple-store-api/models"
	repository "simple-store-api/repositories"

	"gorm.io/gorm"
)

type CategoryUsecase interface {
	Create(category *models.Category) (err error)
	Delete(category *models.Category) (err error)
	GetAll(param *datatransfers.CategoryQueryParams) (categories []*models.Category, cnt int64, err error)
	GetByID(categoryID uint) (category *models.Category, err error)
	Update(category *models.Category) (err error)
}

type categoryUsecase struct {
	db           *gorm.DB
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(db *gorm.DB) CategoryUsecase {
	categoryRepo := repository.NewCategoryRepository(db)
	return &categoryUsecase{
		categoryRepo: categoryRepo,
		db:           db,
	}
}

func (u *categoryUsecase) GetAll(param *datatransfers.CategoryQueryParams) (categories []*models.Category, cnt int64, err error) {
	categories, cnt, err = u.categoryRepo.GetAll(param)
	return
}

func (u *categoryUsecase) GetByID(categoryID uint) (category *models.Category, err error) {
	category, err = u.categoryRepo.GetByID(categoryID)
	return
}

func (u *categoryUsecase) Create(category *models.Category) (err error) {
	err = u.categoryRepo.Create(category, u.db)
	return
}

func (u *categoryUsecase) Update(category *models.Category) (err error) {
	err = u.categoryRepo.Update(category, u.db)
	return
}

func (u *categoryUsecase) Delete(category *models.Category) (err error) {
	err = u.categoryRepo.Delete(category, u.db)
	return
}
