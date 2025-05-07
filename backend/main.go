package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Define your handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Backend!"))
	})

	// Enable CORS with specific settings (allow all origins in this example)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, // Allow frontend URL (adjust accordingly)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Wrap your handler with CORS middleware
	handler := c.Handler(http.DefaultServeMux)

	// Start the server
	fmt.Println("Backend is running on port 8080")
	http.ListenAndServe(":8080", handler)
}
