package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAdminController),
	fx.Provide(NewPlacementController),
)
