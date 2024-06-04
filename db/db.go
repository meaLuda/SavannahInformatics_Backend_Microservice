package db

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


// Global db connections
var DB *gorm.DB

func InitDB(config *Config) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database\n: %v", err)
    }
	// Create tables
    err = DB.AutoMigrate(&Customer{}, &Order{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
}