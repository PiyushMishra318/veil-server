package models

import "gorm.io/gorm"

type Wallet struct {
    gorm.Model  // adds ID, created_at etc.
    OwnerName	string `json:"owner_name"`
    Expenditure	string `json:"expenditure"`
    Savings		string `json:"savings"`
	Phone		string `json:"phone"`
    Balance 	string `json:"balance"`
}