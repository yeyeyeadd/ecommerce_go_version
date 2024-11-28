package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	SellerID    uint    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
