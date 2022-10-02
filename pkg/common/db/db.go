package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"veil/pkg/common/models"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Wallet{})
    db.AutoMigrate(&models.Transaction{})

    return db
}