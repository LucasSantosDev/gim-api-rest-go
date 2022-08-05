package main

import (
	"gim-api-rest-go/database"
	"gim-api-rest-go/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequests()
}