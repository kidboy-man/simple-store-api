package models

import (
	"net/http"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	RoleID       uint           `gorm:"uniqueIndex:role_permission_idx"`
	Role         *Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
	PermissionID uint           `gorm:"uniqueIndex:role_permission_idx"`
	Permission   *Permission    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permission"`
	CreatedAt    time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}

func (r *RolePermission) setAttr() {
	if r.Role != nil {
		r.RoleID = r.Role.ID
	}

	if r.Permission != nil {
		r.PermissionID = r.Permission.ID
	}
}

func (r *RolePermission) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations\
	r.setAttr()
	if r.RoleID == 0 || r.PermissionID == 0 {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_DATA",
		}
		return
	}
	return
}

func (p *RolePermission) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
