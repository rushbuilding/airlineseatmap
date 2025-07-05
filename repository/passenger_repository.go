package repository

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/entity"
)

type PassengerRepository interface {
	Insert(ctx context.Context, data *entity.Passenger)  (int64, error)
	FindById(ctx context.Context, id string) entity.Passenger
} 