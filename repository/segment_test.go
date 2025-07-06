package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/model/entity"
)

var segrepo = NewSegmentManagementRepository(connections.GetConnection())
func TestSegment(t *testing.T) {
	id, _ := uuid.Parse("fd7f5c51-0bc2-417c-a228-9b0ffd97e43c")
	flightId, _ := uuid.Parse("6c3bc909-5c36-4816-a8da-df41530c2e99")
	data := entity.Segment{
		SegmentRef: "3937094243023851424",
		AircraftID: id,
		FlightID: flightId,
		OriginAirportCode: "KUL",
		DestinationAirportCode: "CGK",
		Duration: 2,
		DepartureTime: time.Date(2025, 8, 27, 17, 55, 0, 0, time.UTC),
		DepartureTimezone: "Asia/Jakarta",
		ArrivalTime: time.Date(2025, 8, 27, 17, 55, 0, 0, time.UTC),
		ArrivalTimezone: "Asia/Jakarta",
		DepartureTerminal: "TERMINAL 1",
		ArrivalTerminal: "TERMINAL 2",
		SubjectGovermentApproval: false,
	}
	segrepo.CreateSegment(context.Background(), &data)
	
}

func TestSegmentOffering(t *testing.T) {
	segrepo.CreateSegmentOffering(context.Background(), &entity.SegmentOffering{
		ID: uuid.New(),
		SegmentRef: "3937094243023851424",
		ServiceClass: "ECONOMY",
		FareBasis: "TOWBSSMY",
		FlightMiles: 707,
		AwardFare: false,
		BookingClass: "T",

	})
}

func TestPassengerSegmentOffering(t *testing.T) {
	soId, _ := uuid.Parse("ffc3da1e-273f-4440-a0db-5efd7224a91d")
	pId, _ := uuid.Parse("8fe24407-1df6-4b99-96c9-4dee6b72fdfb")
	segrepo.CreatePassengerSegment(context.Background(), &entity.SegmentPassenger{
		ID: uuid.New(),
		SegmentOfferingID: soId,
		PassengerID: pId,
		PassengerIndex: 1,
		PassengerNameNumber: "01.01",
		MealPreference: "vegetable",
		SeatPreference: "center",
		SeatSelectionEnabled: true,
	})
}

func TestFlight(t *testing.T) {
	segrepo.CreateFlightDetail(context.Background(), &entity.FlightDetail{
		ID: uuid.New(),
		FlightNumber: 302,
		OperatingFlightNumber: 302,
		AirlineCode: "OD",
		OperatingAirlineCode: "OD",
	})
}

func TestGetSegmentOffering(t *testing.T) {
	segrepo.GetSegmentBySegmentRef(context.Background(), "3937094243023851424")
	
}