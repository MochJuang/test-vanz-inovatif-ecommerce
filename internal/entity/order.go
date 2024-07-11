package entity

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:userId"`
	OrderItems []OrderItem `json:"order_items"`
	Total      float64     `gorm:"not null"`
	Status     string      `gorm:"default:pending"`
	CreatedAt  time.Time
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	CreatedAt time.Time
}
