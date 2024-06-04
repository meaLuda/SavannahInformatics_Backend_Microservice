package db

import (
	"time"
	"gorm.io/gorm"
)

// Customer represents a customer in the database
type Customer struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code" gorm:"uniqueIndex:code;type:text;not null"`
}

// Order represents an order in the database
type Order struct {
	gorm.Model
	CustomerID uint      `json:"customer_id"`
	Item       string    `json:"item"`
	Amount     float64   `json:"amount"`
	OrderTime  time.Time `json:"order_time"`
}