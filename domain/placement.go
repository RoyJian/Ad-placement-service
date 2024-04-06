package domain

import "context"

type PlacementRequest struct {
	Gender   string `form:"gender" binding:"omitempty,oneof=M F"`
	Age      int    `form:"age" binding:"omitempty,gte=1,lte=100"`
	Country  string `form:"country" binding:"omitempty,iso3166_1_alpha2"`
	Platform string `form:"platform" binding:"omitempty,oneof=android ios web"`
	Limit    int    `form:"limit,default=5" binding:"gte=1,lte=100"`
	Offset   int    `form:"offset,default=0"`
}
type PlacementResponse struct {
	Item []Advertisement `json:"item"`
}
type PlacementUseCase interface {
	Query(ctx context.Context, params PlacementRequest) ([]Advertisement, error)
}
