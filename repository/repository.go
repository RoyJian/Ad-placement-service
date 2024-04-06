package repository

import "go.uber.org/fx"

var Model = fx.Options(
	fx.Provide(NewAdvertisementRepository))
