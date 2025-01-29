package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading env variables: %v\n", err)
	}
}
func main() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	ADDRESS := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	r.Use(CORSMiddleware())

	// register routes
	r.GET("/", HandleIndex)

	// handle 404 not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	// handle methods not allowed
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	})
	if err := r.Run(ADDRESS); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
