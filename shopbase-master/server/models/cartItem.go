package models

import (
	"time"
)

type CartItem struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserId     uint      `json:"user_id" gorm:"foreignKey"`
	ProductId  string    `json:"product_id" gorm:"foreignKey"`
	Quantity   uint      `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
