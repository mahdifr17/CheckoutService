package database

import (
	"fmt"

	"github.com/mahdifr17/CheckoutService/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(databaseURL, err)
		return nil, err
	}

	// Auto migrate models
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
