package dto

type FlightPassenger struct {
	PassengerIndex  int
	PassengerNameNumber string
	PassengerDetail PassengerDetail
	PassengerInfo PassengerInfo
	Preferences Preferences
	DocumentInfo DocumentInfo
}

type PassengerDetail struct {
	FirstName string
	LastName string
}

type PassengerInfo struct {
	DateOfBirth string
	Gender string
	AgeType string 
	Emails []string
	Phones []string
	Address Address
}

type Address struct {
	Street1 string
	Street2 string
	Postcode string
	State string
	City string
	Country string
	AddressType string

}

type DocumentInfo struct {
	IssuingCountry string
	CountryOfBirth string
	DocumentType string
	Nationality string
}

type Preferences struct {
	SpecialPreferences SpecialPreference
	FrequentFyler []FrequentFyler
}

type SpecialPreference struct {
	MealPreference string
	SeatPreference string
	SpecialRequest []string
	SpecialServiceRequestRemarks []string
}

type FrequentFyler struct {
	Airline string
	Number string
	TierNumber int8
}