package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactGroup struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Active    bool           `json:"active" gorm:"default:true"`
	Contacts  []Contact      `gorm:"many2many:contact_group_members;joinForeignKey:ContactGroupID;JoinReferences:ContactID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
