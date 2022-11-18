package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255)" validate:"required" json:"username"`
	Code      string         `gorm:"index;unique;type:varchar(50)" validate:"required" json:"code"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (Permission) TableName() string {
	return "permission"
}

func (p *Permission) setAttr() {
	p.Code = strings.ToLower(strings.TrimSpace(p.Code))
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	p.setAttr()
	if p.Code == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_DATA",
		}
		return
	}
	return
}

func (p *Permission) BeforeUpdate(tx *gorm.DB) (err error) {
	p.setAttr()
	return
}
