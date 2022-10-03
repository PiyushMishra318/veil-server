package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	OwnerName   string         `json:"owner_name"`
	Expenditure float64        `json:"expenditure"`
	Savings     float64        `json:"savings"`
	Phone       string         `json:"phone"`
	Balance     float64        `json:"balance"`
}
