package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;type:varchar(255)" json:"name"`
	Image       string         `gorm:"type:varchar(255)" json:"value"`
	Description string         `json:"description"`
	Sort        int            `json:"sort"`
	IsActive    *bool          `json:"isActive"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func (Category) TableName() string {
	return "category"
}
