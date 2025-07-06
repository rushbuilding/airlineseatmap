package entity

import "github.com/google/uuid"

type SeatAvailability struct {
	ID                uuid.UUID `gorm:"column:seat_availability_id;primaryKey" json:"id"`
	SegmentRef        string    `gorm:"column:segment_ref" json:"segment_ref"`
	SeatID            uuid.UUID `gorm:"column:seat_id" json:"seat_id"`
	FeeWaived         bool      `gorm:"column:fee_waived" json:"fee_waived"`
	FeeWaiverRuleID   string    `gorm:"column:fee_waiver_rule_id" json:"fee_waiver_rule_id"`
	RefundIndicator   string    `gorm:"column:refund_indicator;size:3" json:"refund_indicator"`
	FreeOfCharge      bool      `gorm:"column:free_of_charge" json:"free_of_charge"`
	IsAvailable       bool      `gorm:"column:is_available" json:"is_available"`

	// Associations
	SeatPrices 		[]SeatPricingOption `gorm:"foreignKey:seat_availability_id;references:seat_availability_id" json:"seat_prices"`
}

func (a *SeatAvailability) TableName() string {
	return "seat_availability"
}