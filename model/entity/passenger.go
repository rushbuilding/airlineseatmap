package entity

import (
	"time"

	"github.com/google/uuid"
)

type Passenger struct {
	PassengerID		uuid.UUID			`gorm:"column:passenger_id;primary_key"`
	FirstName    	string 				`gorm:"column:first_name"`
	LastName     	string				`gorm:"column:last_name"`
	BirthOfDate 	*time.Time 			`gorm:"column:birth_of_date"`
	Gender      	string				`gorm:"column:gender"`
	PassengerType 	string				`gorm:"column:passenger_type"`
	EmailAddresses	[]PassengerEmail	`gorm:"foreignKey:passenger_id;references:passenger_id"`
	PhoneNumbers	[]PassengerPhone	`gorm:"foreignKey:passenger_id;references:passenger_id"`
	Document		PassengerDocument	`gorm:"foreignKey:passenger_id;references:passenger_id"`
}

func (p *Passenger) TableName() string {
	return "passengers"
}