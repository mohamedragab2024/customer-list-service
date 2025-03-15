package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Customer struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Email   string `json:"email"`
}
