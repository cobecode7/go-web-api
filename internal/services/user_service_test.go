package services

import (
    "go-web-api/internal/models"
    "testing"

    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        panic("فشل الاتصال بقاعدة البيانات التجريبية")
    }
    db.AutoMigrate(&models.User{})
    return db
}

func TestUserService_CreateUser(t *testing.T) {
    db := setupTestDB() // ← جديدة
    service := NewUserService(db)

    user := &models.User{
        Name:  "سارة",
        Email: "sara@example.com",
    }

    err := service.CreateUser(user)
    assert.NoError(t, err)

    var savedUser models.User
    result := db.Where("email = ?", user.Email).First(&savedUser)
    assert.NoError(t, result.Error)
    assert.Equal(t, "سارة", savedUser.Name)
}

func TestUserService_UserExistsByEmail(t *testing.T) {
    db := setupTestDB() // ← جديدة
    service := NewUserService(db)

    // أنشئ مستخدمًا
    db.Create(&models.User{Name: "علي", Email: "ali@example.com"})

    exists, err := service.UserExistsByEmail("ali@example.com")
    assert.NoError(t, err)
    assert.True(t, exists)

    exists, err = service.UserExistsByEmail("notexist@example.com")
    assert.NoError(t, err)
    assert.False(t, exists)
}

func TestUserService_GetUsers(t *testing.T) {
    db := setupTestDB() // ← جديدة
    service := NewUserService(db)

    // أنشئ بالضبط مستخدمين
    db.Create(&models.User{Name: "أحمد", Email: "ahmed@example.com"})
    db.Create(&models.User{Name: "منى", Email: "mona@example.com"})

    users, err := service.GetUsers()
    assert.NoError(t, err)
    assert.Len(t, users, 2) // الآن سيكون 2 بالضبط
}
