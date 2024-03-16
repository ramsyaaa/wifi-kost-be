package main

import (
	"log"
	"wifi_kost_be/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables instead")
	}

	db := config.Connect()
	config.Route(db)
}
