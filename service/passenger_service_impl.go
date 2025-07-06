package service

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/repository"
)

type PassengerServiceImpl struct {
	PassengerRepository repository.PassengerRepository
}

func NewPassengerService(PassengerRepository repository.PassengerRepository) PassengerService {
	return &PassengerServiceImpl{PassengerRepository:PassengerRepository}
}

func (r *PassengerServiceImpl) GetPassengerData(ctx context.Context, ref string, passengerId string) (*dto.FlightPassenger, error) {
	passengerData, _ := r.PassengerRepository.FindPassengerDataById(ctx, passengerId)
	passengerSegmentData, err := r.PassengerRepository.FindPassengerSegmentByPassengerIdAndSegmenRef(ctx, passengerId, ref)
	
	if err != nil {
		panic(err)
	}

	var phones []string
	for _, p := range passengerData.PhoneNumbers {
		phones = append(phones, p.PhoneNumber)
	}

	var emails []string
	for _, e := range passengerData.EmailAddresses {
		emails = append(emails, e.EmailAddress)
	}

	address := dto.Address{
		Street1: passengerData.Address.FirstStreetLine,
		Street2: passengerData.Address.SecondStreetLine,
		Postcode: passengerData.Address.PostCode,
		State: passengerData.Address.State,
		City: "",
		Country: passengerData.Address.Country,
		AddressType: passengerData.Address.AddressType,
	}

	passengerDetail := dto.PassengerDetail{
		FirstName: passengerData.FirstName,
		LastName: passengerData.LastName,
	}

	passengerInfo := dto.PassengerInfo{
		DateOfBirth: passengerData.BirthOfDate.Format("2006-01-02"),
		Gender: passengerData.Gender,
		AgeType: passengerData.PassengerType,
		Emails: emails,
		Phones: phones,
		Address: address,
	}

	document := dto.DocumentInfo{
		IssuingCountry: passengerData.Document.IssuingCountry,
		CountryOfBirth: passengerData.Document.CountryOfBirth,
		DocumentType: passengerData.Document.DocumentType,
		Nationality: passengerData.Document.Nationality,
	}

	prefrences := dto.Preferences{
		SpecialPreferences: dto.SpecialPreference{
			MealPreference: passengerSegmentData.MealPreference,
			SeatPreference: passengerSegmentData.SeatPreference,
		},
	}

	pasenger := dto.FlightPassenger{
		PassengerIndex: passengerSegmentData.PassengerIndex,
		PassengerNameNumber: passengerSegmentData.PassengerNameNumber,
		PassengerDetail: passengerDetail,
		PassengerInfo: passengerInfo,
		Preferences: prefrences,
		DocumentInfo: document,
	}

	return &pasenger, nil
}
