package db

import (
	"app/pkg/db/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDBConnection() (*gorm.DB, error) {
	dsn := createDsn(CreateDBConfig())

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database, error: %v", err)
	}

	return db, nil
}

func Migrate() {
	db, err := CreateDBConnection()

	if err != nil {
		log.Fatalf("Failed to migrate database, error: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}

func createDsn(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port)
}
