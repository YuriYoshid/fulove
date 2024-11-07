package handlers

import (
    "net/http"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "github.com/yuri2/fulove/services/user/internal/models"
    "github.com/yuri2/fulove/services/user/internal/database"
)

func SignUp(c *gin.Context) {
    var req models.SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // パスワードのハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
        return
    }

    // ユーザーの作成
    user := models.User{
        Username:      req.Username,
        Email:         req.Email,
        Password:      string(hashedPassword),
        IconURL:       req.IconURL,
        PreferredTime: req.PreferredTime,
        Points:        0,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    // データベースに保存
    result := database.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
        return
    }

    // パスワードを除外してレスポンスを返す
    user.Password = ""
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
        "user": user,
    })
}

func SignIn(c *gin.Context) {
    var req models.SignInRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ユーザーの検索
    var user models.User
    if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // パスワードの検証
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // TODO: JWTトークンの生成（後で実装）
    
    // パスワードを除外してレスポンスを返す
    user.Password = ""
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user": user,
    })
}