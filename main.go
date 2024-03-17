package main

import (
	"Ad_Placement_Service/service/http"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}
	// Init service
	if err := http.Init(); err != nil {
		log.Fatal("Start http service fail!", err)
		return
	}

}
