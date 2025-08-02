package auth

import (
    "go-web-api/internal/services"
    "net/http"

    "github.com/gin-gonic/gin"
		"gorm.io/gorm"
)

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
    Name     string `json:"name" binding:"required,min=2"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "بيانات غير صالحة", "details": err.Error()})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    authService := services.NewAuthService(db)

    err := authService.Register(req.Name, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "فشل التسجيل"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "تم التسجيل بنجاح"})
}

func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "بيانات غير صالحة", "details": err.Error()})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    authService := services.NewAuthService(db)

    token, err := authService.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "بريد إلكتروني أو كلمة مرور غير صحيحة"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
