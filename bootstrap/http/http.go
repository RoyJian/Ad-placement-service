package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
)

type Http interface {
	Start()
	Shutdown(ctx context.Context)
}
type Gin struct {
	Srv *http.Server
	Gin *gin.Engine
}

func NewGin() *Gin {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	addr := fmt.Sprintf("%s:%s", os.Getenv("GIN_HOST"), os.Getenv("GIN_PORT"))
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &Gin{Srv: srv, Gin: router}
}

func (g Gin) Start() {
	if err := g.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Gin service init error ", err)
	}
}

func (g Gin) Shutdown(ctx context.Context) {
	if err := g.Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Print("Close http service success")
}

var Module = fx.Options(fx.Provide(NewGin))
