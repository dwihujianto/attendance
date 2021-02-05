package model

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(191)" json:"name"`
	Email         string         `gorm:"type:varchar(191)" json:"email"`
	Status        bool           `json:"status"`
	LocationPoint string         `gorm:"type:longtext" json:"location_point"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Presences     []Presence
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}
