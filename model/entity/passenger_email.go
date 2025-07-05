package entity

import "github.com/google/uuid"

type PassengerEmail struct {
	EmailID 		uuid.UUID 		`gorm:"column:email_id;primarykey"`
	PassengerID 	uuid.UUID 		`gorm:"column:passenger_id"`
	EmailAddress 	string 			`gorm:"column:email_address"`
}

func (pe *PassengerEmail) TableName() string {
	return "passenger_email"
}