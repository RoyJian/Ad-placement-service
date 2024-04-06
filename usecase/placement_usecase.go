package usecase

import (
	"Ad_Placement_Service/domain"
	"Ad_Placement_Service/repository"
	"context"
)

type PlacementUseCase struct {
	adRepository *repository.AdvertisementRepository
}

func NewPlacementUseCase(adRepository *repository.AdvertisementRepository) *PlacementUseCase {
	return &PlacementUseCase{
		adRepository: adRepository,
	}
}

func (pu *PlacementUseCase) Query(ctx context.Context, params domain.PlacementRequest) ([]domain.Advertisement, error) {
	return pu.adRepository.Query(ctx, params)
}
