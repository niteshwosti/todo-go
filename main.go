package main

import (
	"log"
	"todo/config"
	"todo/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.SetupDB()
	r := routes.SetupRouter()
	r.Run()
}
