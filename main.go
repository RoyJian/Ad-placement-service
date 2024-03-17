package main

import (
	"Ad_Placement_Service/route"
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
	r := route.InitRouter()
	port := fmt.Sprintf(":%s", os.Getenv("GIN_PORT"))
	err = r.Run(port)
	if err != nil {
		log.Fatal("Start backend service fail!")
		return
	}

}
