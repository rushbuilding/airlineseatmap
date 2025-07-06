package repository

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/model/entity"
)

type AircraftManagementRepository interface {
	SetupAircraft(ctx context.Context, aircraft *entity.Aircraft) error
	GetAircraftByID(ctx context.Context, aircraftID string) (*entity.Aircraft, error)
	SetupCabin(ctx context.Context, cabin *entity.Cabin) error
	GetCabinByAircraftIDAndServiceClass(ctx context.Context, aircraftID string, serviceClass string) (*entity.Cabin, error)
	SetupCabinRowAndSeat(ctx context.Context, cabinRow *entity.CabinRow) error
	SetupSeatsPrices(ctx context.Context)
	GetSeatingWithPrice(ctx context.Context, currency string, cabinId string) ([]dto.SeatView, error)
	// GetCabinRowAndSeatByCabinId(ctx context.Context, cabinId string) (*entity.CabinRow, error)
	// SetupSeatDetail(ctx context.Context, seat *entity.SeatDetail) error
	// GetSeatDetailByID(ctx context.Context, seatID string) (*entity.SeatDetail, error)
	// SetupSeatCharacteristic(ctx context.Context, characteristic *entity.SeatCharacteristic) error
	// MapSeatCharacteristic(ctx context.Context, seatCharMap *entity.SeatCharacteristicMap) error
}