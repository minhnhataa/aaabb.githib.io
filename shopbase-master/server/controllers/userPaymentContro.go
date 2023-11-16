package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/database"
	"server/models"
)

func CreateUserPayment(w http.ResponseWriter, r *http.Request) {
	var data models.UserPayment

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parser %v", data)
	}

	database.DB.Where("user_id = ?", data.UserId).Where("type = ?", data.Type).First(&data)

	if data.Id != 0 {
		database.DB.Model(&data).Updates(data)
		fmt.Fprintf(w, "updated user_payment: %v", data)
	} else {
		database.DB.Create(&data)
		fmt.Fprintf(w, "created user_payment: %v", data)
	}
}
