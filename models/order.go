package models

import "time"

type Order struct {
	ID         uint    `gorm:"primaryKey"`
	BuyerID    uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
	CreatedAt  time.Time
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}
