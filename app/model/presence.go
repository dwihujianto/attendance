package model

import(
	"gorm.io/gorm"
	"time"
)

type Presence struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CheckinAt time.Time      `json:"checkin_at`
	CheckoutAt time.Time     `json:"checkout_at`
	CreatedAt time.Time      `json:"created_at`
	UpdatedAt time.Time      `json:"updated_at`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}