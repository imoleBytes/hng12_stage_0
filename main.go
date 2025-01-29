package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env variables: %v\n", err)
	}
}
func main() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	ADDRESS := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	r.GET("/", HandleIndex)

	if err := r.Run(ADDRESS); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
