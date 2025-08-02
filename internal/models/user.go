package models

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"not null" binding:"required,min=2"`
    Email     string    `json:"email" gorm:"uniqueIndex;not null" binding:"required,email"`
    Password  string    `json:"-" gorm:"not null"` // "-" تعني: لا تُرجع في JSON
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
