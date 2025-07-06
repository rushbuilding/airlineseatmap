package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/rushbuilding/airlineseatmap/model/entity"
	"gorm.io/gorm"
)

type SegmentManagementRepositoryImpl struct {
	db *gorm.DB
}

func NewSegmentManagementRepository(db *gorm.DB) SegmentManagementRepository {
	return &SegmentManagementRepositoryImpl{db:db}
}

func (r *SegmentManagementRepositoryImpl) CreateSegment(ctx context.Context, data *entity.Segment) error {
	if data == nil {
		return errors.New("data is required")
	}

	return r.db.Create(data).Error
}

func (r *SegmentManagementRepositoryImpl) CreateSegmentOffering(ctx context.Context, data *entity.SegmentOffering) error {
	if data == nil {
		return errors.New("data is required")
	}
	return r.db.Create(data).Error
}

func (r *SegmentManagementRepositoryImpl) CreateFlightDetail(ctx context.Context, data *entity.FlightDetail) error {
	if data == nil {
		return errors.New("data is required")
	}
	return r.db.Create(data).Error
}

func (r *SegmentManagementRepositoryImpl) GetSegmentOfferingById(ctx context.Context, id string) (*entity.SegmentOffering, error) {
	var segmentOffer entity.SegmentOffering
	err := r.db.Where("segment_offering_id = ?", id).First(&segmentOffer).Error
	if err != nil {
		return nil, err
	}
	return &segmentOffer, nil
}

func (r *SegmentManagementRepositoryImpl) GetSegmentOfferingBySegmentRefAndServiceClass(ctx context.Context, ref string, serviceClass string) (*entity.SegmentOffering, error) {
	var segmentOffer entity.SegmentOffering
	err := r.db.Where("segment_ref = ? AND service_class = ?", ref, serviceClass).First(&segmentOffer).Error
	if err != nil {
		return nil, err
	}
	return &segmentOffer, nil
}

func (r *SegmentManagementRepositoryImpl) GetSegmentBySegmentRef(ctx context.Context, ref string) (*entity.Segment, error) {
	var segment entity.Segment
	err := r.db.Where("segment.segment_ref = ?", ref).
			Joins("FlightDetail").
			Joins("Aircraft").
			Preload("StopOvers").
			First(&segment).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(segment)
	return &segment, nil
}

func (r *SegmentManagementRepositoryImpl) CreatePassengerSegment(ctx context.Context, data *entity.SegmentPassenger) error {
	if data == nil {
		return errors.New("data is required")
	}
	return r.db.Create(data).Error
}

func (r *SegmentManagementRepositoryImpl) GetPassengerSegmentById(ctx context.Context, id string) (*entity.SegmentPassenger, error) {
	var segmentPassenger entity.SegmentPassenger
	
	err := r.db.Where("segment_passenger_id = ?", id).First(&segmentPassenger).Error
	if err != nil {
		return nil, err
	}
	return &segmentPassenger, nil
}