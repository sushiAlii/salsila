package db

import (
	"fmt"
	"os"

	config "github.com/sushiAlii/salsila/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitializeDB() *gorm.DB {
	config.LoadEnv()
	
	dsn := os.Getenv("DB_CONNECTION")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

	return DB
}