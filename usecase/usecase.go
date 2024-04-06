package usecase

import "go.uber.org/fx"

var Model = fx.Options(
	fx.Provide(NewPlacementUseCase),
	fx.Provide(NewAdminUseCase),
)
