package entity

import "github.com/google/uuid"

type PassengerPhone struct {
	ID          uuid.UUID `gorm:"column:phone_id;primaryKey" json:"id"`
	PassengerID uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
}

func (a *PassengerPhone) TableName() string {
	return "passenger_phone"
}