package main

import (
	"fmt"
	"let-you-cook/config"
	"let-you-cook/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.ConfigureLogger()
	config.InitializeMinioClient()
}

func main() {
	router := router.SetupRouter()

	router.Run(":42069")
}
