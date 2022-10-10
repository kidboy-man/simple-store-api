package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CategoryID  uint           `gorm:"index:product_category_idx"`
	Category    *Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Description string         `gorm:"type:text" json:"description"`
	Sort        int            `gorm:"type:integer" json:"sort"`
	IsActive    *bool          `gorm:"type:boolean;default:false" json:"isActive"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`

	Variants []*Variant `json:"variants"`
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) setAttr(tx *gorm.DB) {
	if p.Category != nil {
		p.CategoryID = p.Category.ID
	}

	if p.CategoryID <= 0 {
		tx.Statement.Omits = append(tx.Statement.Omits, "category_id")
	}
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.setAttr(tx)
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.setAttr(tx)
	return
}
