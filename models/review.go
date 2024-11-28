package models

import "time"

type Review struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	OrderID   uint   `gorm:"not null"`
	Rating    int    `gorm:"check:rating BETWEEN 1 AND 5;not null"`
	Comment   string `gorm:"type:text"`
	CreatedAt time.Time
}
