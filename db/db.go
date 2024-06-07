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
    // Check if tables exist and perform auto-migration if they do not
    if !DB.Migrator().HasTable(&Customer{}) {
        log.Println("Customer table does not exist. Creating table...")
        err = DB.AutoMigrate(&Customer{})
        if err != nil {
            log.Fatalf("Error migrating Customer table: %v", err)
        }
    }

    if !DB.Migrator().HasTable(&Order{}) {
        log.Println("Order table does not exist. Creating table...")
        err = DB.AutoMigrate(&Order{})
        if err != nil {
            log.Fatalf("Error migrating Order table: %v", err)
        }
    }

    log.Println("Database initialization complete")
}