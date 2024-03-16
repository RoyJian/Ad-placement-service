package main

import (
	"Ad_Placement_Service/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init service
	r := router.InitRouter()
	port := fmt.Sprintf(":%s", os.Getenv("GIN_PORT"))
	err = r.Run(port)
	if err != nil {
		log.Fatal("Start backend service fail!")
		return
	}

}
