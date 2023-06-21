package models

import (
	"time"
)

type Stock struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Nama      string    `json:"nama"`
	Harga     string    `json:"harga"`
	Qty       string    `json:"qty"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
