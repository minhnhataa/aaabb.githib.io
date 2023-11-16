package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"server/database"
	"server/models"

	"gorm.io/gorm"
)

type Product struct {
	ProductId string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Price     int    `json:"price"`
}

type DataOrder struct {
	Total     int64     `json:"total"`
	PaymentId uint      `json:"payment_id"`
	Products  []Product `json:"products"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	database.DB.Transaction(func(tx *gorm.DB) error {
		// check dang nhap
		cookie, err := r.Cookie("user_id")
		if err != nil {
			fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
			return nil
		}

		// parse string to uint
		user_id, err := strconv.ParseUint(cookie.Value, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Cannot parse string to uint: %v", err)
			return nil
		}

		// parser
		var data DataOrder
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Fprintf(w, "error when body parser %v", data)
			return nil
		}

		// check total
		total := 0
		for i := 0; i < len(data.Products); i++ {
			total = total + int(data.Products[i].Quantity)*data.Products[i].Price
		}

		if int64(total) != data.Total {
			fmt.Fprintf(w, "khong khop so tien tong cac san pham")
			return nil
		}

		order := models.Order{
			UserId:     uint(user_id),
			Total:      data.Total,
			PaymentId:  data.PaymentId,
			Status:     true,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}

		if err := tx.Create(&order).Error; err != nil {
			fmt.Fprintf(w, "cannot create order")
			return err
		}

		for _, v := range data.Products {
			order_detail := models.OrderDetail{
				ProductId:  v.ProductId,
				Quantity:   v.Quantity,
				OrderId:    order.Id,
				CreatedAt:  time.Now(),
				ModifiedAt: time.Now(),
			}
			if err := tx.Create(&order_detail).Error; err != nil {
				fmt.Fprintf(w, "cannot create order detail")
				return err
			}
		}

		fmt.Fprintf(w, "order thanh cong: %v", order)
		return nil
	})
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	// check dang nhap
	cookie, err := r.Cookie("user_id")
	if err != nil {
		fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
		return
	}

	user_id, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Cannot parse string to uint: %v", err)
	}

	var orders []models.Order
	database.DB.Where("user_id = ?", user_id).Find(&orders)
	json.NewEncoder(w).Encode(orders)
}
