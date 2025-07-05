package entity

import "github.com/google/uuid"

type FlightDetail struct {
	FlightID 				uuid.UUID 	`gorm:"column:flight_id;primarykey"`
	FlightNumber 			int16 		`gorm:"column:flight_number"`
	OperatingFlightNumber 	int16 		`gorm:"column:operating_flight_number"`
	AirlineCode 			string 		`gorm:"column:airline_code"`
	OperatingAirlineCode 	string 		`gorm:"column:operationg_airline_code"`
}

func (f *FlightDetail) TableName() string {
	return "flight_detail"
}