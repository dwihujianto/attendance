package model

import (
	"gorm.io/gorm"
	"time"
)

type Presence struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	EmployeeId         uint           `json:"employee_id"`
	CheckinAt          time.Time      `json:"checkin_at`
	CheckoutAt         time.Time      `json:"checkout_at`
	CheckinPoint       string         `gorm:"longtext" json:"checkin_point"`
	CheckoutPoint      string         `gorm:"longtext" json:"checkout_point"`
	CheckinNote        string         `gorm:"text" json:"checkin_note"`
	CheckoutNote       string         `gorm:"text" json:"checkout_note"`
	CheckinAttachment  string         `gorm:"varchar(191)" json:"checkin_attachment"`
	CheckoutAttachment string         `gorm:"varchar(191)" json:"checkout_attachment"`
	CreatedAt          time.Time      `json:"created_at`
	UpdatedAt          time.Time      `json:"updated_at`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
