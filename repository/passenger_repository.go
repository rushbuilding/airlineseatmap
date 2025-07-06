package repository

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/entity"
)

type PassengerRepository interface {
	InsertPassengerData(ctx context.Context, data *entity.Passenger) error
	FindPassengerDataById(ctx context.Context, id string) (*entity.Passenger, error)

	FindPassengerSegmentByPassengerIdAndSegmenRef(ctx context.Context, pid string, ref string) (*entity.SegmentPassenger, error)
} 