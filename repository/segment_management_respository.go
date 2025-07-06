package repository

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/entity"
)

type SegmentManagementRepository interface {
	CreateSegment(ctx context.Context, data *entity.Segment) error
	CreateSegmentOffering(ctx context.Context, data *entity.SegmentOffering) error
	CreateFlightDetail(ctx context.Context, data *entity.FlightDetail) error
	CreatePassengerSegment(ctx context.Context, data *entity.SegmentPassenger) error

	GetSegmentOfferingById(ctx context.Context, id string) (*entity.SegmentOffering, error)
	GetSegmentOfferingBySegmentRefAndServiceClass(ctx context.Context, ref string, serviceClass string) (*entity.SegmentOffering, error)
	GetSegmentBySegmentRef(ctx context.Context, ref string) (*entity.Segment, error)
	
}