package routes

import (
    "go-web-api/internal/db"
    "go-web-api/internal/handlers"
    "go-web-api/internal/handlers/auth"
    "go-web-api/internal/middleware"

    "github.com/gin-gonic/gin"
    "net/http" // ✅ تم إضافته
)

func Setup(r *gin.Engine) {
    // أضف قاعدة البيانات إلى السياق
    r.Use(func(c *gin.Context) {
        c.Set("db", db.DB)
        c.Next()
    })

    api := r.Group("/api")
    {
        // مسارات عامة
        api.POST("/register", auth.Register)
        api.POST("/login", auth.Login)

        // مسارات محمية
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.GET("/profile", func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{"message": "أنت مسجل دخول"})
            })
            protected.GET("/users", handlers.GetUsers)
            protected.POST("/users", handlers.CreateUser)
        }
    }
}
