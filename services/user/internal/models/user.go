package models

import (
    "time"
)

type User struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Email       string    `json:"email" gorm:"uniqueIndex;not null"`  // uniqueIndexを追加
    Password    string    `json:"password" gorm:"not null"`           // jsonタグを修正
    Username    string    `json:"username" gorm:"uniqueIndex"`        // uniqueIndexを追加
    IconURL     string    `json:"icon_url"`
    BathTimes   []BathTime `json:"bath_times" gorm:"foreignKey:UserID"`
    Points      int       `json:"points" gorm:"default:0"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// 入浴記録
type BathRecord struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    BathTime    time.Time `json:"bath_time"`
    Duration    int       `json:"duration"`     // 入浴時間（分）
    Temperature float32   `json:"temperature"`  // 温度
    Points      int       `json:"points"`       // 獲得ポイント
    Comment     string    `json:"comment"`      // コメント
    CreatedAt   time.Time `json:"created_at"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
}

// 入浴予定時間
type BathTime struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    DayOfWeek   string    `json:"day_of_week"` // "monday", "tuesday", etc.
    Time        string    `json:"time"`        // "HH:MM" format
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type SignUpRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type SignInRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type ProfileUpdateRequest struct {
    Username  string            `json:"username" binding:"required"`
    IconURL   string           `json:"icon_url"`
    BathTimes map[string]string `json:"bath_times"` // {"monday": "21:00", "tuesday": "22:00", ...}
}

type AuthResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}