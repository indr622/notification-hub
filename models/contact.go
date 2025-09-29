package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name          string         `json:"name" gorm:"not null"`
	Email         string         `json:"email" gorm:"unique;not null"`
	Phone         string         `json:"phone" gorm:"unique;not null"`
	UserType      string         `json:"user_type" gorm:"not null"`
	ContactSource string         `json:"contact_source" gorm:"not null"`
	Active        bool           `json:"active" gorm:"default:true"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
