package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"server/database"
	"server/models"

	"github.com/gorilla/mux"
)

func CreateCartItem(w http.ResponseWriter, r *http.Request) {
	var data models.CartItem
	cookie, err := r.Cookie("user_id")
	if err != nil {
		fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
		return
	}
	user_id, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Cannot parse string to uint: %v", err)
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}
	var cartitem models.CartItem
	data.UserId = uint(user_id)

	database.DB.Where("user_id = ?", user_id).Where("product_id = ?", data.ProductId).First(&cartitem)

	if cartitem.Id != 0 {
		cartitem.Quantity = cartitem.Quantity + data.Quantity
		database.DB.Updates(&cartitem)
		fmt.Fprintf(w, "updated cartitem of user %v", data.UserId)
		return
	}

	data.CreatedAt = time.Now()
	data.ModifiedAt = time.Now()

	database.DB.Create(&data)
	fmt.Fprintf(w, "created cartitem of user %v", user_id)
}

func GetCartItems(w http.ResponseWriter, r *http.Request) {

	type CartCustom struct {
		Id        int    `json:"id"`
		ProductId string `json:"product_id"`
		Quantity  uint   `json:"quantity"`
		Name      string `json:"name"`
		Price     string `json:"price"`
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
		return
	}
	user_id, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Cannot parse string to uint: %v", err)
	}

	var cartitems models.CartItem
	var cartCustom []CartCustom
	database.DB.Model(&cartitems).Select("cart_items.id, product_id, quantity, name, price").Joins("LEFT JOIN products ON products.id = product_id").Where("user_id = ?", uint(user_id)).Find(&cartCustom)

	fmt.Println(cartCustom)
	json.NewEncoder(w).Encode(cartCustom)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user_id")
	if err != nil {
		fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.CartItem

	database.DB.Where("id = ?", id).First(&data)
	if data.Id == 0 {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "err when parse body")
		return
	}

	data.ModifiedAt = time.Now()

	database.DB.Model(&data).Updates(data)

	fmt.Fprintf(w, "updated cartitem_id: %v", data.Id)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	data := models.CartItem{
		Id: uint(id),
	}

	database.DB.Delete(&data)

	fmt.Fprintf(w, "deleted cartitem_id %v", data.Id)
}
