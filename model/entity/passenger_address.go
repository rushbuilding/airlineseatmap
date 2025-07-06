package entity

import "github.com/google/uuid"

type PassengerAddress struct {
	ID               uuid.UUID `gorm:"column:address_id;primaryKey" json:"id"`
	PassengerID      uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	FirstStreetLine  string    `gorm:"column:first_street_line" json:"first_street_line"`
	SecondStreetLine string    `gorm:"column:second_street_line" json:"second_street_line"`
	PostCode         string    `gorm:"column:postcode" json:"postcode"`
	State            string    `gorm:"column:state" json:"state"`
	Country          string    `gorm:"column:country" json:"country"`
	AddressType      string    `gorm:"column:address_type" json:"address_type"`
}

func (a *PassengerAddress) TableName() string {
	return "passenger_address"
}