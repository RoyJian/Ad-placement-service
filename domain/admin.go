package domain

import "context"

type Admin struct {
	Advertisement Advertisement
}

type AdminUseCase interface {
	Create(ctx context.Context, admin *Admin) error
}
