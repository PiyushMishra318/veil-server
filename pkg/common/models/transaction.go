package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	WalletID  uuid.UUID      `json:"wallet_id"`
	Amount    float64        `json:"amount"`
	// type is either credit card, debit card or cash
	Type         string    `json:"type"`
	Recurring    bool      `json:"recurring"`
	RecurIntCnt  float64   `json:"recur_int_count"`
	RecurIntPer  string    `json:"recur_int_per"`
	CategoryID   uuid.UUID `json:"category_id"`
	Category     Category  `json:"category"`
	VoiceFile    string    `json:"voice_file"`
	VoiceMessage string    `json:"vocie_message"`
}
