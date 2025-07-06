package entity

import "github.com/google/uuid"

type SegmentPassenger struct {
	ID                  uuid.UUID `gorm:"column:segment_passenger_id;primaryKey" json:"id"`
	SegmentOfferingID   uuid.UUID `gorm:"column:segment_offering_id" json:"segment_offering_id"`
	PassengerID         uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	SeatAvailabilityID  *string `gorm:"column:seat_availability_id" json:"seat_availability_id"` // optional
	PassengerIndex      int    `gorm:"column:passenger_index" json:"passenger_index"`
	PassengerNameNumber string `gorm:"column:passenger_name_number" json:"passenger_name_number"`
	SeatSelectionEnabled bool  `gorm:"column:seat_selection_enabled" json:"seat_selection_enabled"`
	MealPreference      string `gorm:"column:meal_preference" json:"meal_preference"`
	SeatPreference      string `gorm:"column:seat_preference" json:"seat_preference"`

	//association
	SegmentOffering    SegmentOffering   `gorm:"foreignKey:segment_offering_id;references:segment_offering_id" json:"segment_offering"`
	SeatAvailability   *SeatAvailability `gorm:"foreignKey:seat_availability_id;references:seat_availability_id" json:"seat_availability"`
	SpecialRequests 	[]SpecialRequest `gorm:"foreignKey:segment_passenger_id;references:segment_passenger_id" json:"special_requests"`
}

func (a *SegmentPassenger) TableName() string {
	return "segment_passenger"
}

