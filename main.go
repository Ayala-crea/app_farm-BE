package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gobizdevelop/config"
	"gobizdevelop/routes"

	"github.com/rs/cors"
)

func main() {
	// Connect to MongoDB
	connectdb := config.Mongoconn
	if config.ErrorMongoconn != nil {
		fmt.Println("Failed to connect to MongoDB:", config.ErrorMongoconn)
		return
	}

	// Check if the connection is successful
	if connectdb != nil {
		fmt.Println("Successfully connected to MongoDB!")
	} else {
		fmt.Println("MongoDB connection is nil")
	}

	postgresDB := config.PostgresConn
	if config.ErrorPostgresConn != nil {
		fmt.Println("Failed to connect to PostgreSQL:", config.ErrorPostgresConn)
		return
	}

	// Check if the PostgreSQL connection is successful
	if postgresDB != nil {
		fmt.Println("Successfully connected to PostgreSQL!")
	} else {
		fmt.Println("PostgreSQL connection is nil")
	}

	// Initialize the router from the routes package
	router := routes.InitializeRoutes()

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
