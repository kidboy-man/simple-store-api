package models

import (
	"time"

	"gorm.io/gorm"
)

type Price struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	VariantID uint
	Variant   *Variant       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"variant"`
	Nominal   float32        `gorm:"type:NUMERIC" json:"nominal"`
	Currency  string         `gorm:"type:varchar(3)" json:"currency"`
	IsActive  *bool          `gorm:"type:boolean" json:"isActive"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (Price) TableName() string {
	return "variant"
}
