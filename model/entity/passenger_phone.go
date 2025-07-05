package entity

import "github.com/google/uuid"

type PassengerPhone struct {
	PhoneID 		uuid.UUID	 	`gorm:"column:phone_id;primarykey"`
	PassengerID 	uuid.UUID	 	`gorm:"column:passenger_id"`
	PhoneNumber		string			`gorm:"column:phone_number"`
}

func (pp *PassengerPhone) TableName() string {
	return "passenger_phone"
}