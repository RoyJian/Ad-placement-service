package routes

import (
	"go.uber.org/fx"
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

func NewRoute(adRoute *AdvertisementRoute, healthRoute *HealthRoute) *Routes {
	return &Routes{adRoute, healthRoute}
}

var Module = fx.Options(
	fx.Provide(NewRoute),
	fx.Provide(NewAdvertisementRoute),
	fx.Provide(NewHealthRoute),
)

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
