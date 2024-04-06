package bootstrap

import (
	"Ad_Placement_Service/bootstrap/db"
	"Ad_Placement_Service/bootstrap/http"
	"Ad_Placement_Service/routes"
	"context"
	"go.uber.org/fx"
	"log"
)

func Run(
	gin *http.Gin,
	db *db.MongoDb,
	route *routes.Routes,
	lc fx.Lifecycle,
) {

	lc.Append(fx.Hook{
		OnStart: func(_ctx context.Context) error {
			go func() {
				route.Setup()
				gin.Start()
			}()
			log.Println("Server start At ", gin.Srv.Addr, "🚀🚀🚀")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			gin.Shutdown(ctx)
			db.Disconnect(ctx)
			return nil
		},
	})
}
