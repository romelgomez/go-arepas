package main

import (
	"fmt"
	"go-arepas/config"
	"go-arepas/helper"
	"go-arepas/router"

	// repositories
	post_repository "go-arepas/domain/post/repository"

	// services
	post_service "go-arepas/domain/post/service"

	// controllers
	post_controller "go-arepas/domain/post/controller"

	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}

	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	// repositories
	postRepository := post_repository.NewPostRepository(db)

	// services
	postService := post_service.NewPostServiceImpl(postRepository)

	// controllers
	postController := post_controller.NewPostController(postService)

	// router
	routes := router.NewRouter(
		postController,
	)

	// Read CORS settings from environment variables
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	allowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")

	// Wrap your routes with CORS handling middleware using the environment variable values
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedHeaders(allowedHeaders),
	)(routes)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        corsHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running on port %s\n", port)
	server_err := server.ListenAndServe()

	if server_err != nil {
		panic(server_err)
	}
}
