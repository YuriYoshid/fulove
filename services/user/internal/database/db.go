package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/yuri2/fulove/services/user/internal/models"
)

var DB *gorm.DB

func InitDB() {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbUser, dbPassword, dbName)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    err = db.AutoMigrate(&models.User{}, &models.BathRecord{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    DB = db
    log.Println("Database connection established")
}

func GetDB() *gorm.DB {
    return DB
}