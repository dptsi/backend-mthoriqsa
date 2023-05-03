// models/customer.go
package models

import (
	"time"
)

type Customer struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"cust_name"`
	Phone     string    `json:"cust_phone"`
	DOB       time.Time `json:"cust_dob"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
