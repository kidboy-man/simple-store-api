package models

import (
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	ProductID   uint
	Product     *Product       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Description string         `gorm:"type:text" json:"description"`
	Sort        int            `gorm:"type:integer" json:"sort"`
	Stock       int            `gorm:"type:integer" json:"stock"`
	IsActive    *bool          `gorm:"type:boolean" json:"isActive"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func (Variant) TableName() string {
	return "variant"
}
