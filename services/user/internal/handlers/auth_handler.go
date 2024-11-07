package handlers

import (
    "errors"
    "log"
    "net/http"
    "time"
	"strconv"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/yuri2/fulove/services/user/internal/models"
    "github.com/yuri2/fulove/services/user/internal/database"
)

func SignUp(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    var req models.SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Failed to bind request: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // メールアドレスの重複チェック
    var existingUser models.User
    if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
        return
    } else if !errors.Is(err, gorm.ErrRecordNotFound) {
        log.Printf("Database error during email check: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    // パスワードのハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Password hashing failed: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // 新規ユーザーの作成
    user := models.User{
        Email:     req.Email,
        Password:  string(hashedPassword),
        Points:    0,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // データベースに保存
    result := database.DB.Create(&user)
    if result.Error != nil {
        log.Printf("User creation failed: %v", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // レスポンスからパスワードを除外
    user.Password = ""
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
        "user": user,
    })
}

func SignIn(c *gin.Context) {
    var req models.SignInRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Failed to bind signin request: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // ユーザーの検索
    var user models.User
    result := database.DB.Where("email = ?", req.Email).First(&user)
    if result.Error != nil {
        log.Printf("User not found: %v", result.Error)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // パスワードの検証
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        log.Printf("Password verification failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // 入浴時間情報の取得
    var bathTimes []models.BathTime
    if err := database.DB.Where("user_id = ?", user.ID).Find(&bathTimes).Error; err != nil {
        log.Printf("Failed to fetch bath times: %v", err)
    }
    user.BathTimes = bathTimes

    // パスワードを除外してレスポンスを返す
    user.Password = ""
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user": user,
    })
}

// UpdateProfile handler を追加
func UpdateProfile(c *gin.Context) {
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 32)
    if err != nil {
        log.Printf("Invalid user ID: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req models.ProfileUpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Failed to bind profile update request: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // トランザクション開始
    tx := database.DB.Begin()

    // ユーザー情報の更新
    if err := tx.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
        "username":   req.Username,
        "icon_url":   req.IconURL,
        "updated_at": time.Now(),
    }).Error; err != nil {
        tx.Rollback()
        log.Printf("Failed to update user profile: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    // 既存の入浴時間を削除
    if err := tx.Where("user_id = ?", userID).Delete(&models.BathTime{}).Error; err != nil {
        tx.Rollback()
        log.Printf("Failed to delete existing bath times: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bath times"})
        return
    }

    // 新しい入浴時間を登録
    for day, timeStr := range req.BathTimes {
        bathTime := models.BathTime{
            UserID:    uint(userID),
            DayOfWeek: day,
            Time:      timeStr,
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }
        if err := tx.Create(&bathTime).Error; err != nil {
            tx.Rollback()
            log.Printf("Failed to create new bath time: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bath times"})
            return
        }
    }

    // トランザクションをコミット
    if err := tx.Commit().Error; err != nil {
        log.Printf("Failed to commit transaction: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save changes"})
        return
    }

    // 更新されたユーザー情報を取得
    var updatedUser models.User
    if err := database.DB.Preload("BathTimes").First(&updatedUser, userID).Error; err != nil {
        log.Printf("Failed to fetch updated user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Profile updated successfully",
        "user": updatedUser,
    })
}