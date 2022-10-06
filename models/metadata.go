package models

import (
	"time"

	"gorm.io/gorm"
)

type Metadata struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Type      string         `gorm:"unique;type:varchar(255)" json:"type"`
	Value     string         `json:"value"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (Metadata) TableName() string {
	return "metadata"
}
