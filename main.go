package main

import (
	"ecommerce-api/models"
	"ecommerce-api/routes"
	"log"
)

func main() {
	// Initialize database
	models.InitDB()

	// Initialize router
	r := routes.InitRoutes()

	// Run server
	log.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
