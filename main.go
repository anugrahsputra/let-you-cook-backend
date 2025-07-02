package main

import (
	"fmt"
	"let-you-cook/config"
	"let-you-cook/router"
	minio_util "let-you-cook/utils/minio"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.ConfigureLogger()
	minio_util.InitMinio()
}

func main() {
	router := router.SetupRouter()

	router.Run(":42069")
}
