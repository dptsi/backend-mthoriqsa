// models/customer.go
package models

import (
	"time"
)

type Customer struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Nama      string    `json:"nama"`
	Telepon   string    `json:"telepon"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
