package entity

import (
	"time"

	"github.com/google/uuid"
)

type Segment struct {
	SegmentRef              string    `gorm:"column:segment_ref;primaryKey" json:"segment_ref"`
	AircraftID				uuid.UUID `gorm:"column:aircraft_id" json:"aircraft_id"`
	FlightID				uuid.UUID `gorm:"column:flight_id" json:"flight_id"`
	OriginAirportCode       string    `gorm:"column:origin_airport_code" json:"origin_airport_code"`
	DestinationAirportCode  string    `gorm:"column:destination_airport_code" json:"destination_airport_code"`
	DepartureTime           time.Time `gorm:"column:departure_time" json:"departure_time"`
	DepartureTimezone       string    `gorm:"column:departure_timezone" json:"departure_timezone"`
	ArrivalTime             time.Time `gorm:"column:arrival_time" json:"arrival_time"`
	ArrivalTimezone         string    `gorm:"column:arrival_timezone" json:"arrival_timezone"`
	Duration                int16     `gorm:"column:duration" json:"duration"`
	DepartureTerminal       string    `gorm:"column:departure_terminal" json:"departure_terminal"`
	ArrivalTerminal         string    `gorm:"column:arrival_terminal" json:"arrival_terminal"`
	SubjectGovermentApproval bool     `gorm:"column:subject_to_goverment_approval" json:"subject_government_approval"`

	//Association
	FlightDetail			FlightDetail 	`gorm:"foreignKey:flight_id;references:flight_id" json:"flight_detail"`
	Aircraft 				Aircraft 		`gorm:"foreignKey:aircraft_id;references:aircraft_id" json:"aircraft"`
	StopOvers				[]SegmentStopOver `gorm:"foreignKey:segment_ref;references:segment_ref" json:"stop_overs"`
}

func (a *Segment) TableName() string {
	return "segment"
}