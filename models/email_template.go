package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmailTemplate struct {
	ID           string         `json:"id" gorm:"type:uuid;primaryKey"`
	TemplateName string         `json:"template_name" gorm:"unique;not null"`
	EmailTitle   string         `json:"email_title"`
	EmailBody    string         `json:"email_body" gorm:"type:text"`
	PropertyInfo string         `json:"property_info"`
	SenderName   string         `json:"sender_name"`
	Signature    string         `json:"signature"`
	Active       bool           `json:"active" gorm:"default:true"`
	Html         bool           `json:"html" gorm:"default:true"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// Hook untuk generate UUID otomatis sebelum create
func (e *EmailTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}
	return
}
