package entity

import "github.com/google/uuid"

type SegmentOffering struct {
	ID            uuid.UUID `gorm:"column:segment_offering_id;primaryKey" json:"id"`
	SegmentRef    string `gorm:"column:segment_ref" json:"segment_ref"`
	ServiceClass  string `gorm:"column:service_class" json:"service_class"` // ENUM
	FareBasis     string `gorm:"column:fare_basis" json:"fare_basis"`
	FlightMiles   int    `gorm:"column:flight_miles" json:"flight_miles"`
	AwardFare     bool   `gorm:"column:award_fare" json:"award_fare"`
	BookingClass  string `gorm:"column:booking_class;size:2" json:"booking_class"`
}

func (a *SegmentOffering) TableName() string {
	return "segment_offering"
}
