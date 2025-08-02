package services

import (
    "go-web-api/internal/models"
    "gorm.io/gorm"
)

type UserService struct {
    DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{DB: db}
}

func (s *UserService) GetUsers() ([]models.User, error) {
    var users []models.User
    if err := s.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.DB.Create(user).Error
}

func (s *UserService) UserExistsByEmail(email string) (bool, error) {
    var count int64
    err := s.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
