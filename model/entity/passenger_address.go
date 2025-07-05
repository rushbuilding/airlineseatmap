package entity

import "github.com/google/uuid"

type PassengerAddress struct {
	AddressID 			uuid.UUID 		`gorm:"column:address_id;primarykey"`
	PassengerID 		uuid.UUID 		`gorm:"column:passenger_id"`
	FirstStreetLine 	string 			`gorm:"column:first_street_line"`
	SecondStreetLine 	string			`gorm:"column:second_street_line"`
	PostCode 			string			`gorm:"column:postcode"`
	State				string			`gorm:"column:state"`
	Country				string 			`gorm:"column:country"`
	AddressType			string 			`gorm:"column:address_type"`
}

func (pd *PassengerAddress) TableName() string {
	return "passenger_address"
}