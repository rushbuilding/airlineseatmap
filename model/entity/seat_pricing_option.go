package entity

import "github.com/google/uuid"

type SeatPricingOption struct {
	ID                uuid.UUID  `gorm:"column:pricing_id;primaryKey" json:"id"`
	SeatAvailabilityID uuid.UUID `gorm:"column:seat_availability_id" json:"seat_availability_id"`
	Currency          string  `gorm:"column:currency" json:"currency"`
	PriceAmount       float64 `gorm:"column:price_amount" json:"price_amount"`
	TaxAmount         float64 `gorm:"column:tax_amount" json:"tax_amount"`
}

func (a *SeatPricingOption) TableName() string {
	return "seat_pricing_option"
}