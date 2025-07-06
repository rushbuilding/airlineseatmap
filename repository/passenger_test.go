package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/model/entity"
)


var repository = NewPassengerRepository(connections.GetConnection())
func TestInsertData(t *testing.T) {
	ctx := context.Background()

	passenger_id := uuid.New()

	birthDate := time.Date(1970, time.August, 17, 0, 0, 0, 0, time.UTC)

	passenger := entity.Passenger{
		ID: passenger_id,
		FirstName: "Rutwik",
		LastName: "Sabre",
		BirthOfDate: &birthDate,
		Gender: "MALE",
		PassengerType: "ADT",
		EmailAddresses: []entity.PassengerEmail {
			entity.PassengerEmail{
				ID: uuid.New(),
				PassengerID: passenger_id,
				EmailAddress: "johnsabre@domain.com",
			},
		},
		PhoneNumbers: []entity.PassengerPhone {
			entity.PassengerPhone{
				ID: uuid.New(),
				PassengerID: passenger_id,
				PhoneNumber: "08123456789",
			},
		},
		Document: entity.PassengerDocument{
			ID: uuid.New(),
			PassengerID: passenger_id,
			DocumentType: "P",
			IssuingCountry: "ID",
			CountryOfBirth: "ID",
			Nationality: "ID",
		},
			
	}

	err := repository.InsertPassengerData(ctx, &passenger)
	if err != nil {
		panic(err)
	}
}

func TestGetData(t *testing.T) {
	ctx := context.Background()
	repository.FindPassengerDataById(ctx, "8fe24407-1df6-4b99-96c9-4dee6b72fdfb")
}

func TestPassengerData(t *testing.T) {
	repository.FindPassengerSegmentByPassengerIdAndSegmenRef(context.Background(), "8fe24407-1df6-4b99-96c9-4dee6b72fdfb", "3937094243023851424")
}