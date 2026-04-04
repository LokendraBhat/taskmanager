package main

import (
	"crud-api/internal/config"
	"crud-api/internal/handlers"
	"crud-api/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	config.InitDB()
	defer config.DB.Close()

	// Insert default user if env vars are set
	config.InitDefaultUser()

	// Routes
	http.HandleFunc("/", handlers.LoginPage)
	http.HandleFunc("/login", middleware.CSRFMiddleware(handlers.LoginAction))
	http.HandleFunc("/logout", handlers.LogoutAction)
	http.HandleFunc("/dashboard", middleware.AuthMiddleware(handlers.DashboardPage))
	http.HandleFunc("/task/create", middleware.AuthMiddleware(middleware.CSRFMiddleware(handlers.CreateTask)))
	http.HandleFunc("/task/update", middleware.AuthMiddleware(middleware.CSRFMiddleware(handlers.UpdateTask)))
	http.HandleFunc("/task/toggle", middleware.AuthMiddleware(middleware.CSRFMiddleware(handlers.ToggleTask)))
	http.HandleFunc("/task/delete", middleware.AuthMiddleware(middleware.CSRFMiddleware(handlers.DeleteTask)))

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}