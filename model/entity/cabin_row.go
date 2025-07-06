package entity

import "github.com/google/uuid"

type CabinRow struct {
	ID               uuid.UUID  	`gorm:"column:row_id;primaryKey" json:"id"`
	CabinID          uuid.UUID  	`gorm:"column:cabin_id" json:"cabin_id"`
	RowNumber        int     		`gorm:"column:row_number" json:"row_number"`
	IsEnable         bool    		`gorm:"column:is_enable" json:"is_enable"`
	DisableCause     *string 		`gorm:"column:disable_cause" json:"disable_cause"`
	RowType          *string 		`gorm:"column:row_type" json:"row_type"`                   // EXIT, WING_BEGIN, WING_END
	RowDesignation   *string 		`gorm:"column:row_designation" json:"row_designation"`     // FRONT_OF_CABIN, VACANT_OR_OFFERED_LAST
	
	// Associations
	Seats 			[]SeatDetail 	`gorm:"foreignKey:ID;references:ID" json:"seats"`
}

func (a *CabinRow) TableName() string {
	return "cabin_row"
}