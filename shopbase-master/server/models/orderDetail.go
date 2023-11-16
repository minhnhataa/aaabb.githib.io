package models

import "time"

type OrderDetail struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	ProductId  string    `json:"product_id"`
	Quantity   uint      `json:"quantity"`
	OrderId    uint      `json:"order_id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
