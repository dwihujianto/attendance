package model

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at`
	UpdatedAt time.Time      `json:"updated_at`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"type:varchar(191);unique" json:"name"`
	City      string         `gorm:"type:varchar(191)" json:"city"`
	Age       int            `json:"age"`
	Status    bool           `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}
