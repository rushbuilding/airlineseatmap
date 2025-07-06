package entity

import "github.com/google/uuid"

type PassengerDocument struct {
	ID             uuid.UUID `gorm:"column:document_id;primaryKey" json:"id"`
	PassengerID    uuid.UUID `gorm:"column:passenger_id" json:"passenger_id"`
	DocumentType   string    `gorm:"column:document_type" json:"document_type"`
	IssuingCountry string    `gorm:"column:issuing_country" json:"issuing_country"`
	CountryOfBirth string    `gorm:"column:country_of_birth" json:"country_of_birth"`
	Nationality    string    `gorm:"column:nationality" json:"nationality"`
}

func (a *PassengerDocument) TableName() string {
	return "passenger_document"
}