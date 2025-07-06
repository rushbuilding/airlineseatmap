package service

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/dto"
)

type SegmentService interface {
	GetSegmentData(ctx context.Context, ref string, class string) (*dto.FlightSegment, error)
}