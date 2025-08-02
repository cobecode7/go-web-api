package main

import (
    "go-web-api/internal/db"
    "go-web-api/internal/routes"
    "go-web-api/internal/models"
    "log"
    "os"

    "github.com/gin-gonic/gin"
		_ "github.com/joho/godotenv/autoload" // ← يحمل .env تلقائيًا عند التشغيل
)

func main() {
		 log.Println("JWT_SECRET =", os.Getenv("JWT_SECRET"))
    // الاتصال بقاعدة البيانات
    if err := db.Connect(); err != nil {
        log.Fatal("فشل الاتصال بقاعدة البيانات:", err)
    }
		// الهجرة: إنشاء الجدول إذا لم يكن موجودًا
		if err := db.DB.AutoMigrate(&models.User{}); err != nil {
    		log.Fatal("فشل الهجرة (AutoMigrate):", err)
		}

    // إنشاء مثيل Gin
    r := gin.Default()

    // CORS (للتطوير مع frontend مثل React)
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // تحميل المسارات
    routes.Setup(r)

    // تشغيل الخادم على PORT من البيئة
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("🚀 الخادم يعمل على المنفذ :%s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("فشل تشغيل الخادم:", err)
    }
}
