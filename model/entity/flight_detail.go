package entity

import "github.com/google/uuid"

type FlightDetail struct {
	ID              	  uuid.UUID `gorm:"column:flight_id;primaryKey" json:"flight_id"`
	FlightNumber          int16     `gorm:"column:flight_number" json:"flight_number"`
	OperatingFlightNumber int16     `gorm:"column:operating_flight_number" json:"operating_flight_number"`
	AirlineCode           string    `gorm:"column:airline_code" json:"airline_code"`
	OperatingAirlineCode  string    `gorm:"column:operating_airline_code" json:"operating_airline_code"`
}

func (a *FlightDetail) TableName() string {
	return "flight_detail"
}