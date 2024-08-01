package main

import (
	api "github.com/alejandrososa/mars-rover-go/internal/app/adapters/api/http"

	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := api.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
