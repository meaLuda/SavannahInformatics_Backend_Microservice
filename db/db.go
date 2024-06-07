package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


// Global db connections
var DB *gorm.DB

func InitDB(config *Config) {
    dsn := "host=" + config.DBHost +
           " user=" + config.DBUser +
           " dbname=" + config.DBName +
           " password=" + config.DBPassword +
           " port=" + config.DBPort +
           " sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database\n: %v", err)
    }
	// Create tables
    // err = DB.AutoMigrate(&Customer{}, &Order{})
    // if err != nil {
    //     log.Fatalf("Error migrating database: %v", err)
    // }
}