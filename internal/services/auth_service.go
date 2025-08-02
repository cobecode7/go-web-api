package services

import (
		"errors"
    "go-web-api/internal/models"
    "golang.org/x/crypto/bcrypt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "gorm.io/gorm"
)

type AuthService struct {
    DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{DB: db}
}

// HashPassword يُشفّر كلمة المرور
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash يتحقق من تطابق كلمة المرور
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateJWT يُولّد توكن JWT
func GenerateJWT(email string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return "", errors.New("متغير البيئة JWT_SECRET غير معرّف")
    }

    claims := &jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // صالح لمدة 24 ساعة
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

// Register يسجّل مستخدمًا جديدًا
func (s *AuthService) Register(name, email, password string) error {
    if _, err := s.DB.DB(); err != nil {
        return err
    }

    // التحقق من وجود المستخدم
    var user models.User
    if s.DB.Where("email = ?", email).First(&user).RowsAffected > 0 {
        return nil // يمكن تحسينه ليرجع خطأ
    }

    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }

    user = models.User{
        Name:     name,
        Email:    email,
        Password: hashedPassword,
    }

    return s.DB.Create(&user).Error
}

// Login يتحقق من بيانات الدخول ويُرجع توكن
func (s *AuthService) Login(email, password string) (string, error) {
    var user models.User
    if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return "", errors.New("المستخدم غير موجود")
    }

    if !CheckPasswordHash(password, user.Password) {
        return "", errors.New("كلمة المرور غير صحيحة")
    }

    return GenerateJWT(user.Email)
}
