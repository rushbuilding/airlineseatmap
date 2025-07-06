package service

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/dto"
)

type PassengerService interface {
	GetPassengerData(ctx context.Context, ref string, passengerId string) (*dto.FlightPassenger, error)
}