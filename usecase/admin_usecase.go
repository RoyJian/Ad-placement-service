package usecase

import (
	"Ad_Placement_Service/domain"
	"Ad_Placement_Service/repository"
	"context"
)

type AdminUseCase struct {
	adRepository *repository.AdvertisementRepository
}

func (a AdminUseCase) Create(ctx context.Context, admin *domain.Admin) error {
	return a.adRepository.Create(ctx, &admin.Advertisement)
}

func NewAdminUseCase(adRepository *repository.AdvertisementRepository) *AdminUseCase {
	return &AdminUseCase{adRepository: adRepository}
}
