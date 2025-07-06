package entity

import "github.com/google/uuid"

type FrequentFlyer struct {
	ID               uuid.UUID `gorm:"column:frequent_flyer_id;primaryKey" json:"id"`
	PassengerID      uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	AirlineCode      string    `gorm:"column:airline_code" json:"airline_code"`
	MembershipNumber string    `gorm:"column:membership_number" json:"membership_number"`
	TierNumber       int       `gorm:"column:tier_number" json:"tier_number"`
}

func (a *FrequentFlyer) TableName() string {
	return "frequent_flyer"
}