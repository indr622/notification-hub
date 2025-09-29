package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID                uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GroupName         string         `json:"group_name" gorm:"not null"`
	NotificationEvent string         `json:"notification_event"`
	Site              string         `json:"site"`
	Description       string         `json:"description"`
	Channels          string         `json:"channels"`
	Contacts          string         `json:"contacts"`
	ContactGroups     string         `json:"contact_groups"`
	Active            bool           `json:"active" gorm:"default:true"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}
