package entity

import "github.com/google/uuid"

type Aircraft struct {
	AircraftID 				uuid.UUID 	`gorm:"column:aircraft_id;primarykey"`
	RegistrationNumber 		int16 		`gorm:"column:registration_number"`
	RegistrationName 		int16 		`gorm:"column:registration_name"`
}

func (a *Aircraft) TableName() string {
	return "aircraft"
}