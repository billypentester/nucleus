package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system envs")
	}
	log.Println("- env file loaded from: .env")
}
