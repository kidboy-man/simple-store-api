package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"index;unique;type:varchar(255)" validate:"required" json:"username"`
	Password  string    `gorm:"type:varchar(255)" validate:"required" json:"-"`
	Email     string    `gorm:"index;unique;type:varchar(255)" validate:"required,email" json:"email"`
	Token     string    `gorm:"-" json:"token"`
	CreatedAt time.Time `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) setAttr() {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	if u.Username == "" || u.Password == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_AUTH_DATA",
		}
		return
	}
	u.setAttr()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.setAttr()
	return
}
