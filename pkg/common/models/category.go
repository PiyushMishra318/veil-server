package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID             uuid.UUID `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `json:"name"`
	ParentCategory string         `json:"parent_category" gorm:"default:null"`
	Show           bool           `json:"show"`
	Image          string         `json:"image"`
}
