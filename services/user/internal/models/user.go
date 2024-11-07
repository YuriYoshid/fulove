package models

import (
    "time"
)

type User struct {
    ID              uint      `json:"id" gorm:"primaryKey"`
    Username        string    `json:"username" gorm:"unique;not null"`
    Email           string    `json:"email" gorm:"unique;not null"`
    Password        string    `json:"-" gorm:"not null"`  // JSONでは非表示
    IconURL         string    `json:"icon_url"`
    PreferredTime   string    `json:"preferred_time" gorm:"type:time"` // 希望入浴時間
    Points          int       `json:"points" gorm:"default:0"`
    LastBathTime    time.Time `json:"last_bath_time"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

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

type SignUpRequest struct {
    Username        string `json:"username" binding:"required"`
    Email           string `json:"email" binding:"required,email"`
    Password        string `json:"password" binding:"required,min=6"`
    IconURL         string `json:"icon_url"`
    PreferredTime   string `json:"preferred_time" binding:"required"`
}

type SignInRequest struct {
    Email       string `json:"email" binding:"required,email"`
    Password    string `json:"password" binding:"required"`
}

type AuthResponse struct {
    Token    string `json:"token"`
    User     User   `json:"user"`
}