package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/database"
	"server/models"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	id := data.Id

	var product models.Product
	database.DB.Where("id = ?", id).First(&product)

	if product.Id != "" {
		fmt.Fprintf(w, "product_id exists")
		return
	}

	database.DB.Create(&data)
	fmt.Fprintf(w, "created product %v", data)
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.Product

	database.DB.Where("id = ?", id).First(&data)
	if data.Id == "" {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.Product

	database.DB.Where("id = ?", id).First(&data)
	if data.Id == "" {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "err when parse body")
		return
	}

	database.DB.Model(&data).Updates(data)

	fmt.Fprintf(w, "updated product_id: %v", data.Id)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data := models.Product{
		Id: id,
	}
	database.DB.Delete(&data)

	fmt.Fprintf(w, "deleted product_id: %v", id)
}
