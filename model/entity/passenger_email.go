package entity

import "github.com/google/uuid"

type PassengerEmail struct {
	ID           uuid.UUID `gorm:"column:email_id;primaryKey" json:"id"`
	PassengerID  uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	EmailAddress string    `gorm:"column:email_address" json:"email_address"`
}

func (a *PassengerEmail) TableName() string {
	return "passenger_email"
}