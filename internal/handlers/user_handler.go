package handlers

import (
    "go-web-api/internal/models"
    "net/http"

    "github.com/gin-gonic/gin"
		"gorm.io/gorm"
		"go-web-api/internal/services"
)

func GetUsers(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    service := services.NewUserService(db)

    users, err := service.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "فشل جلب المستخدمين"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "بيانات غير صالحة", "details": err.Error()})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    service := services.NewUserService(db)

    exists, err := service.UserExistsByEmail(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "خطأ في التحقق من البريد"})
        return
    }
    if exists {
        c.JSON(http.StatusConflict, gin.H{"error": "البريد الإلكتروني مسجل مسبقًا"})
        return
    }

    if err := service.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "فشل إنشاء المستخدم"})
        return
    }

    c.JSON(http.StatusCreated, user)
}
