package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() *gorm.DB {
	dsn := "postgresql://postgres:december181996@localhost:5432/salsila?sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

	return DB
}