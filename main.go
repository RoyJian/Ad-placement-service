package main

import (
	"Ad_Placement_Service/bootstrap"
	"Ad_Placement_Service/bootstrap/cache"
	"Ad_Placement_Service/bootstrap/db"
	"Ad_Placement_Service/bootstrap/http"
	"Ad_Placement_Service/controllers"
	"Ad_Placement_Service/repository"
	"Ad_Placement_Service/routes"
	"Ad_Placement_Service/usecase"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}
	app := fx.New(
		fx.NopLogger,
		db.Module,
		cache.Module,
		routes.Module,
		http.Module,
		controllers.Module,
		usecase.Model,
		repository.Model,
		fx.Invoke(bootstrap.Run),
	)
	app.Run()

}
