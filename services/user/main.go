package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/yuri2/fulove/services/user/internal/database"
    "github.com/yuri2/fulove/services/user/internal/handlers"
)

func main() {
    database.InitDB()

    r := gin.Default()

    // CORSミドルウェア
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Writer.Header().Set("Access-Control-Max-Age", "3600")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // 静的ファイルの提供
    r.Static("/images", "./uploads/images")

    // ヘルスチェック
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // ユーザー関連エンドポイント
    r.POST("/signup", handlers.SignUp)
    r.POST("/signin", handlers.SignIn)
    r.PUT("/users/:id/profile", handlers.UpdateProfile)

    // 画像アップロードエンドポイント
    r.POST("/upload/image", handlers.UploadImage)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}