package models

import (
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ProductID   uint           `gorm:"index"`
	Product     *Product       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Description string         `gorm:"type:text" json:"description"`
	Sort        int            `gorm:"type:integer" json:"sort"`
	Stock       int            `gorm:"type:integer" json:"stock"`
	IsActive    *bool          `gorm:"type:boolean;default:false" json:"isActive"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`

	Prices []*Price `json:"prices"`
}

func (Variant) TableName() string {
	return "variant"
}

func (v *Variant) setAttr(tx *gorm.DB) {
	if v.Product != nil {
		v.ProductID = v.Product.ID
	}

	if v.ProductID <= 0 {
		tx.Statement.Omits = append(tx.Statement.Omits, "product_id")
	}
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	v.setAttr(tx)
	return
}

func (v *Variant) BeforeUpdate(tx *gorm.DB) (err error) {
	v.setAttr(tx)
	return
}
