package service

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/dto"
)

type CabinService interface {
	GetSeatingMapAndPriceByCabinId(ctx context.Context, segmentRef string, serviceClass string, currency string) (dto.SeatMap, error)
}