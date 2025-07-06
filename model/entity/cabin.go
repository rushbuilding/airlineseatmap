package entity

import "github.com/google/uuid"

type Cabin struct {
	ID              uuid.UUID `gorm:"column:cabin_id;primaryKey" json:"id"`
	AircraftID      string 	`gorm:"column:aircraft_id" json:"aircraft_id"`
	ServiceClass    string `gorm:"column:service_class" json:"service_class"` // ECONOMY, BUSINESS, FIRST_CLASS
	Deck            string `gorm:"column:deck;default:MAIN" json:"deck"`       // MAIN, UPPER
	FirstRowNumber  int    `gorm:"column:first_row_number" json:"first_row_number"`
	LastRowNumber   int    `gorm:"column:last_row_number" json:"last_row_number"`

	// Associations
	// CabinRows        []CabinRow           `gorm:"foreignKey:cabin_id;references:cabin_id" json:"cabin_rows"`
	CabinLayouts     []CabinConfiguration `gorm:"foreignKey:cabin_id;references:cabin_id" json:"cabin_layouts"`
}

func (a *Cabin) TableName() string {
	return "cabin"
}