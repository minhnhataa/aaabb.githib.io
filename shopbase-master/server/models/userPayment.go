package models

type UserPayment struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Type string `json:"type"`
	UserId uint `json:"user_id" gorm:"foreignKey"`
}