package model

import (
	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name   string `gorm:"type:varchar(191);unique" json:"name"`
	City   string `gorm:"type:varchar(191)" json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}
