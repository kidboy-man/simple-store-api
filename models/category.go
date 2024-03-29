package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Description string         `gorm:"type:text" json:"description"`
	Sort        *int           `gorm:"type:integer" json:"sort"`
	IsActive    *bool          `gorm:"type:boolean;default:false" json:"isActive"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`

	Products []*Product `json:"products"`
}

func (Category) TableName() string {
	return "category"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if it passes through here, we are lacking of validations
	if c.Name == "" || c.Image == "" || c.Description == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "MISSING_FIELD_VALUE_CREATE_CATEGORY",
		}
		return
	}
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
