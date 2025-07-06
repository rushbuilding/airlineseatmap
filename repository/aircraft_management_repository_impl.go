package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/model/entity"
	"gorm.io/gorm"
)

type AircraftManagementRepositoryImpl struct {
	db *gorm.DB
}

func NewAircraftManagementRepositoryImpl(db *gorm.DB) AircraftManagementRepository {
	return &AircraftManagementRepositoryImpl{db:db}
}

func (r *AircraftManagementRepositoryImpl) SetupAircraft(ctx context.Context, data *entity.Aircraft) error {
	if data == nil {
		return errors.New("data is required")
	}

	return r.db.Create(data).Error
}

func (r *AircraftManagementRepositoryImpl) GetAircraftByID(ctx context.Context, aircraftID string)  (*entity.Aircraft, error) {
	if aircraftID == "" {
		return nil, errors.New("data is required")
	}

	var aircraft entity.Aircraft 
	err := r.db.Where("aircraft_id = ?", aircraftID).First(&aircraft).Error
	if err != nil {
		return nil, err
	}

	return &aircraft, nil
}

func (r *AircraftManagementRepositoryImpl) SetupCabin(ctx context.Context, data *entity.Cabin) error {
	if data == nil {
		return errors.New("data is required")
	}

	return r.db.Create(data).Error
}

func (r *AircraftManagementRepositoryImpl) GetCabinByAircraftIDAndServiceClass(ctx context.Context, aircraftID string, serviceClass string) (*entity.Cabin, error) {
	if aircraftID == "" || serviceClass == "" {
		return nil, errors.New("required aircraft id & service class as paramater")
	}

	var cabin entity.Cabin
	err := r.db.Where("cabin.aircraft_id = ? AND cabin.service_class = ?", aircraftID, serviceClass).Preload("CabinLayouts", func(db *gorm.DB) *gorm.DB {
        return db.Order("column_number ASC")
    }).First(&cabin).Error
	if err != nil {
		return nil, err
	}

	return &cabin, err
}

func (r *AircraftManagementRepositoryImpl) SetupCabinRowAndSeat(ctx context.Context, data *entity.CabinRow) error {
	if data == nil {
		return errors.New("data is required")
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Seats").Create(&data).Error; err != nil {
			return err
		}

		if len(data.Seats) > 0 {
			if err := tx.CreateInBatches(&data.Seats, 3).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (r *AircraftManagementRepositoryImpl) GetSeatingWithPrice(ctx context.Context, currency string, cabinId string) ([]dto.SeatView, error) {
	var seatview []dto.SeatView
	err := r.db.Raw(`SELECT 
		sd.store_front_slot_code,
		sd.is_selectable,
		sa.fee_waived,
		sa.fee_waiver_rule_id,
		sa.refund_indicator,
		sa.free_of_charge,
		sa.is_available,
		spo.currency,
		spo.price_amount,
		spo.tax_amount,
		spo.price_amount + spo.tax_amount as total,
		cr.row_number,
		cr.is_enable,
		cr.disable_cause,
		cr.row_type,
		cr.row_designation,
		cc.column_code
	FROM seat_detail sd
	JOIN seat_availability sa ON sd.seat_id = sa.seat_id
	JOIN seat_pricing_option spo ON sa.seat_availability_id = spo.seat_availability_id
	JOIN cabin_row cr ON sd.row_id = cr.row_id
	JOIN cabin_configuration cc ON sd.cabin_layout_id = cc.cabin_layout_id
	WHERE sd.cabin_id = ?
	AND spo.currency = ?
	ORDER BY cr.row_number, cc.column_number;`, cabinId, currency).Scan(&seatview).Error
	if err != nil {
		return nil, err
	}

	return seatview, nil

}

func (r *AircraftManagementRepositoryImpl) SetupSeatsPrices(ctx context.Context) {
	var seats []entity.SeatDetail

	// Step 1: Get seat details
	if err := r.db.Where("cabin_id = ?", "69bf771a-b203-43da-b179-7c0ac3039386").Find(&seats).Error; err != nil {
		panic(err)
	}

	// Step 2: Prepare seat availability data with nested prices
	var seatAvailability []entity.SeatAvailability
	for _, s := range seats {
		availabilityID := uuid.New()
		availability := entity.SeatAvailability{
			ID:             availabilityID,
			SegmentRef:     "3937094243023851424",
			SeatID:         s.ID,
			FeeWaived:      false,
			RefundIndicator: "R",
			FreeOfCharge:   false,
			IsAvailable:    true,
			SeatPrices: []entity.SeatPricingOption{
				{
					ID:                 uuid.New(),
					SeatAvailabilityID: availabilityID,
					Currency:           "MYR",
					PriceAmount:        30.0,
					TaxAmount:          0.0,
				},
			},
		}
		seatAvailability = append(seatAvailability, availability)
	}

	// Optional: print to debug
	if _, err := json.MarshalIndent(seatAvailability, "", "  "); err == nil {
		// fmt.Println(string(jsonData))
	}

	// Step 3: Save atomically using transaction
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, avail := range seatAvailability {
			// Insert SeatAvailability
			if err := tx.Create(&avail).Error; err != nil {
				return fmt.Errorf("failed to insert SeatAvailability: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}