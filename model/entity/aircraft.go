package entity

import "github.com/google/uuid"

type Aircraft struct {
	ID                 uuid.UUID `gorm:"column:aircraft_id;primaryKey" json:"id"`
	RegistrationNumber int16     `gorm:"column:registration_number" json:"registration_number"`
	RegistrationName   string    `gorm:"column:registration_name" json:"registration_name"`
}

func (a *Aircraft) TableName() string {
	return "aircraft"
}