package main

import (
	"server/database"
	"server/routes"
)

func main() {
	database.Connect()
	routes.Init()
}
