package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	CategoryID  uint
	Category    *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Description string         `gorm:"type:text" json:"description"`
	Sort        int            `gorm:"type:integer" json:"sort"`
	IsActive    *bool          `gorm:"type:boolean" json:"isActive"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`

	Variants []*Variant `json:"variants"`
}

func (Product) TableName() string {
	return "product"
}
