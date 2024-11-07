package handlers

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
    // アップロードディレクトリの作成
    uploadDir := "uploads/images"
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        log.Printf("Failed to create upload directory: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
        return
    }

    // ファイルの取得
    file, err := c.FormFile("image")
    if err != nil {
        log.Printf("Failed to get file: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    // ファイル名の生成
    ext := filepath.Ext(file.Filename)
    newFileName := uuid.New().String() + ext
    filePath := filepath.Join(uploadDir, newFileName)

    // ファイルの保存
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        log.Printf("Failed to save file: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    // URLの生成
    imageURL := fmt.Sprintf("/images/%s", newFileName)

    c.JSON(http.StatusOK, gin.H{
        "url": imageURL,
        "message": "File uploaded successfully",
    })
}