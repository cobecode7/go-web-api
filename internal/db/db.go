package db

import (
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Println("⚠️ DATABASE_URL غير معرّف، استخدام الإعدادات الافتراضية")
        dsn = "host=localhost user=go_user password=go_password_123 dbname=go_web_db port=5432 sslmode=disable"
    }

    var err error
    maxRetries := 10
    for i := 0; i < maxRetries; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("✅ الاتصال بقاعدة البيانات نجح")
            return nil
        }
        log.Printf("❌ فشل الاتصال بقاعدة البيانات (المحاولة %d/%d): %v", i+1, maxRetries, err)
        time.Sleep(3 * time.Second)
    }

    return err
}
