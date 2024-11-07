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

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    // 認証関連のエンドポイント
    r.POST("/signup", handlers.SignUp)
    r.POST("/signin", handlers.SignIn)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}