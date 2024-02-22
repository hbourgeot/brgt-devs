package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Brgt struct {
	DB *gorm.DB
}

func NewConnection() (*Brgt, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Brgt{DB: db}, nil
}
