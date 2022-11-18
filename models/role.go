package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255)" validate:"required" json:"username"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (Role) TableName() string {
	return "role"
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	if r.Name == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_DATA",
		}
		return
	}
	return
}

func (p *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
