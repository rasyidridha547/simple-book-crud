package models

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var sqlDB *sql.DB

func ConnectDatabase() {
	destination := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&Book{})

	DB = db
}

func CloseDatabase() error {
	if sqlDB != nil {
		return sqlDB.Close()
	}

	return nil
}
