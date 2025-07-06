package entity

import (
	"time"

	"github.com/google/uuid"
)

type Passenger struct {
	ID              uuid.UUID           `gorm:"column:passenger_id;primaryKey" json:"id"`
	FirstName       string              `gorm:"column:first_name" json:"first_name"`
	LastName        string              `gorm:"column:last_name" json:"last_name"`
	BirthOfDate     *time.Time          `gorm:"column:birth_of_date" json:"birth_of_date"`
	Gender          string              `gorm:"column:gender" json:"gender"`
	PassengerType   string              `gorm:"column:passenger_type" json:"passenger_type"`
	
	// Associations
	Address 		PassengerAddress 	`gorm:"foreignKey:passenger_id;references:passenger_id" json:"address"`
	EmailAddresses  []PassengerEmail    `gorm:"foreignKey:passenger_id;references:passenger_id" json:"email_addresses"`
	PhoneNumbers    []PassengerPhone    `gorm:"foreignKey:passenger_id;references:passenger_id" json:"phone_numbers"`
	Document        PassengerDocument   `gorm:"foreignKey:passenger_id;references:passenger_id" json:"document"`
	FrequentFlyers	[]FrequentFlyer		`gorm:"foreignKey:passenger_id;references:passenger_id" json:"frequent_flyers"`
}

func (a *Passenger) TableName() string {
	return "passengers"
}