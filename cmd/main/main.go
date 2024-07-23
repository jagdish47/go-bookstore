package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jagdish47/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Initialize a new router
	r := mux.NewRouter()

	// Register bookstore routes
	routes.RegisterBookStoreRoutes(r)

	// Handle all routes with the router
	http.Handle("/", r)

	// Start the server on localhost:9010 and log any errors
	log.Println("Starting server on :9010")
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
