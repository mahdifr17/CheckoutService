package database

import (
	"fmt"

	"github.com/mahdifr17/CheckoutService/internals/infra/repository/postgresql"
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

	db.AutoMigrate(
		postgresql.Order{},
		postgresql.Product{},
		postgresql.Promotion{},
		postgresql.OrderItem{},
	)

	return db, nil
}
