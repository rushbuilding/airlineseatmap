package entity

import "github.com/google/uuid"

type PassengerDocument struct {
	DocumentID 		uuid.UUID 		`gorm:"column:document_id;primarykey"`
	PassengerID 	uuid.UUID 		`gorm:"column:passenger_id"`
	DocumentType 	string 			`gorm:"column:document_type"`
	IssuingCountry 	string			`gorm:"column:issuing_country"`
	CountryOfBirth 	string			`gorm:"column:country_of_birth"`
	Nationality		string			`gorm:"column:nationality"`
}

func (pd *PassengerDocument) TableName() string {
	return "passenger_document"
}