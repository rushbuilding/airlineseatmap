package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rushbuilding/airlineseatmap/model/entity"
	"gorm.io/gorm"
)

type PassengerRepositoryImpl struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) PassengerRepository {
	return &PassengerRepositoryImpl{db:db}
}

func (r *PassengerRepositoryImpl) InsertPassengerData(ctx context.Context, data *entity.Passenger) error {
	if data == nil {
		return errors.New("data is required")
	}

	return r.db.WithContext(ctx).Create(data).Error
}

func (r *PassengerRepositoryImpl) FindPassengerDataById(ctx context.Context, id string) (*entity.Passenger, error) {
	var passenger entity.Passenger
	err := r.db.WithContext(ctx).Where("passengers.passenger_id = ?", id).
				Preload("EmailAddresses").
				Preload("PhoneNumbers").
				Joins("Document").
				First(&passenger).Error

	return &passenger, err
}

func (r *PassengerRepositoryImpl) FindPassengerSegmentByPassengerIdAndSegmenRef(ctx context.Context, id string, ref string) (*entity.SegmentPassenger, error) {
	var segmentPassenger entity.SegmentPassenger
	err := r.db.Model(&entity.SegmentPassenger{}).
	Joins("JOIN segment_offering ON segment_offering.segment_offering_id = segment_passenger.segment_offering_id").
	Preload("SegmentOffering").
	Where("segment_passenger.passenger_id = ? AND segment_offering.segment_ref = ?", id, ref).
	First(&segmentPassenger).Error

	jsonData, _ := json.MarshalIndent(segmentPassenger, "", "  ")
	fmt.Println(string(jsonData))

	return &segmentPassenger, err
}