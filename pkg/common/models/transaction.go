package models

import "gorm.io/gorm"

type Transaction struct {
    gorm.Model  // adds ID, created_at etc.
}