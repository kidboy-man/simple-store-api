package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	RoleID    uint           `gorm:"index:admin_role_idx"`
	Role      *Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
	Name      string         `gorm:"index;unique;type:varchar(255)" validate:"required" json:"username"`
	Password  string         `gorm:"type:varchar(255)" validate:"required" json:"-"`
	Email     string         `gorm:"index;unique;type:varchar(255)" validate:"required,email" json:"email"`
	Token     string         `gorm:"-" json:"token"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (Admin) TableName() string {
	return "admin"
}

func (a *Admin) setAttr() {
	a.Email = strings.ToLower(strings.TrimSpace(a.Email))
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	a.setAttr()
	if a.Email == "" || a.Password == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_AUTH_DATA",
		}
		return
	}

	return
}

func (a *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	a.setAttr()
	return
}
