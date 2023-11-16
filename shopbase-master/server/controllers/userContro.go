package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/database"
	"server/models"
)

type Data struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var data Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	if data.Password != data.Confirm {
		fmt.Fprintf(w, "password do not match")
		return
	}

	var isExist models.User

	database.DB.Where("email = ?", data.Email).First(&isExist)

	if isExist.Email != "" {
		fmt.Fprintf(w, "email exists")
		return
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	database.DB.Create(&user)
	fmt.Fprintf(w, "create user: %v", user)
}
