package http

import (
	"Ad_Placement_Service/routes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var srv *http.Server

func Init(ctx context.Context) error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.RegisterRouter(router)
	addr := fmt.Sprintf("%s:%s", os.Getenv("GIN_HOST"), os.Getenv("GIN_PORT"))
	srv = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
func Shutdown(ctx context.Context) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Print("Close http service success")
}
