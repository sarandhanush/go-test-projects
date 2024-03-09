package main

import (
	"fmt"
	"go-test-projects/login/handler"
	"net/http"
)

func init() {
	fmt.Println("Initializing...........")
}

func main() {
	// Register the login handler
	http.HandleFunc("/login", handler.LoginHandler)

	// Start the web server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
