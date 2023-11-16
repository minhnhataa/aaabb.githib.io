package routes

import (
	"fmt"
	"log"
	"net/http"

	"server/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Welcome).Methods("GET")

	// auth
	router.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	// product
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controllers.AllProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// cart_item
	router.HandleFunc("/api/cartitems", controllers.CreateCartItem).Methods("POST")
	router.HandleFunc("/api/cartitems", controllers.GetCartItems).Methods("GET")
	router.HandleFunc("/api/cartitems/{id}", controllers.UpdateCartItem).Methods("PUT")
	router.HandleFunc("/api/cartitems/{id}", controllers.DeleteCartItem).Methods("DELETE")

	// user_payment
	router.HandleFunc("/api/userpayments", controllers.CreateUserPayment).Methods("POST")

	// order
	router.HandleFunc("/api/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders", controllers.GetAllOrders).Methods("GET")

	// catalog
	// router.HandleFunc("/api/catalog", controllers.CreateCatalog).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	fmt.Println("server is running http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
