package main

import (
	"Ad_Placement_Service/service/cache"
	"Ad_Placement_Service/service/http"
	"Ad_Placement_Service/service/mongodb"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}
	go func() {
		// Init mongoDB service
		if err := mongodb.Init(ctx); err != nil {
			log.Fatal("start mongoDB service fail ", err)
		}
		// Init Redis service
		if err := cache.Init(ctx); err != nil {
			log.Fatal("start redis service fail! ", err)
		}
		// Init gin service
		if err := http.Init(ctx); err != nil {
			log.Fatal("Start http service fail! ", err)
		}

	}()
	defer mongodb.Disconnect(ctx)
	defer http.Shutdown(ctx)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server exiting...")

}
