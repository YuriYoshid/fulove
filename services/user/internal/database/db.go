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
    // データベース接続情報
    dbConfig := struct {
        Host     string
        User     string
        Password string
        DBName   string
        Port     string
    }{
        Host:     getEnv("DB_HOST", "localhost"),
        User:     getEnv("DB_USER", "fulove"),
        Password: getEnv("DB_PASSWORD", "fulove_password"),
        DBName:   getEnv("DB_NAME", "fulove_db"),
        Port:     getEnv("DB_PORT", "5432"),
    }

    // 接続文字列の構築
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
        dbConfig.Host,
        dbConfig.User,
        dbConfig.Password,
        dbConfig.DBName,
        dbConfig.Port,
    )

    // データベースに接続
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // マイグレーションの実行（テーブルの作成）
    err = DB.AutoMigrate(&models.User{}, &models.BathTime{}, &models.BathRecord{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Database connection established and migrations completed")
}

// 環境変数を取得（デフォルト値付き）
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}