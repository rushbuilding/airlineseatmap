package repository

import (
	"context"
	"fmt"
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

	birthDate := time.Date(1998, time.January, 15, 0, 0, 0, 0, time.UTC)

	passenger := entity.Passenger{
		PassengerID: passenger_id,
		FirstName: "Alex",
		LastName: "Alex",
		BirthOfDate: &birthDate,
		Gender: "MALE",
		PassengerType: "ADT",
		EmailAddresses: []entity.PassengerEmail {
			entity.PassengerEmail{
				EmailID: uuid.New(),
				PassengerID: passenger_id,
				EmailAddress: "alexalex@gmail.com",
			},
		},
		PhoneNumbers: []entity.PassengerPhone {
			entity.PassengerPhone{
				PhoneID: uuid.New(),
				PassengerID: passenger_id,
				PhoneNumber: "08123456789",
			},
		},
		Document: entity.PassengerDocument{
			DocumentID: uuid.New(),
			PassengerID: passenger_id,
			DocumentType: "P",
			IssuingCountry: "ID",
			CountryOfBirth: "ID",
			Nationality: "ID",
		},	
	}

	res, err := repository.Insert(ctx, &passenger)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetData(t *testing.T) {
	ctx := context.Background()
	repository.FindById(ctx, "8fe24407-1df6-4b99-96c9-4dee6b72fdfb")
}