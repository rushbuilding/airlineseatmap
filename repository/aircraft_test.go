package repository

import (
	"context"
	"encoding/json"

	// "encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/rushbuilding/airlineseatmap/db/connections"
	"github.com/rushbuilding/airlineseatmap/model/entity"
	// "golang.org/x/tools/go/analysis/passes/appends"
)

var repo = NewAircraftManagementRepositoryImpl(connections.GetConnection())

func TestAircraftInset(t *testing.T) {
	data := entity.Aircraft{
		ID:                 uuid.New(),
		RegistrationNumber: 703,
		RegistrationName:   "Boeing 737",
	}

	repo.SetupAircraft(context.Background(), &data)
}

func TestGetAircraft(t *testing.T) {
	res, err := repo.GetAircraftByID(context.Background(), "fd7f5c51-0bc2-417c-a228-9b0ffd97e43c")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestSetupCabin(t *testing.T) {
	cabinId := uuid.New()
	err := repo.SetupCabin(context.Background(), &entity.Cabin{
		ID:             cabinId,
		AircraftID:     "fd7f5c51-0bc2-417c-a228-9b0ffd97e43c",
		ServiceClass:   "ECONOMY",
		Deck:           "MAIN",
		FirstRowNumber: 4,
		LastRowNumber:  35,
		CabinLayouts: []entity.CabinConfiguration{
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "LEFT_SIDE",
				ColumnNumber: 1,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "A",
				ColumnNumber: 2,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "B",
				ColumnNumber: 3,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "C",
				ColumnNumber: 4,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "AISLE",
				ColumnNumber: 5,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "D",
				ColumnNumber: 6,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "E",
				ColumnNumber: 7,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "F",
				ColumnNumber: 8,
			},
			{
				ID:           uuid.New(),
				CabinID:      cabinId,
				ColumnCode:   "RIGHT_SIDE",
				ColumnNumber: 9,
			},
		},
	})
	if err != nil {
		panic(err)
	}

}

func TestCabinByAircraftIDServiceClass(t *testing.T) {
	res, err := repo.GetCabinByAircraftIDAndServiceClass(context.Background(), "fd7f5c51-0bc2-417c-a228-9b0ffd97e43c", "ECONOMY")
	if err != nil {
		panic(err)
	}
	jsonData, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonData))

}

func TestCabinRowSetup(t *testing.T) {

	res, _ := repo.GetCabinByAircraftIDAndServiceClass(context.Background(),
		"fd7f5c51-0bc2-417c-a228-9b0ffd97e43c", "ECONOMY")

	for rowNum := res.FirstRowNumber; rowNum <= res.LastRowNumber; rowNum++ {
		rowID := uuid.New()
		isWing := rowNum >= 10 && rowNum < 20

		var rowDesignation *string
		switch rowNum {
		case res.FirstRowNumber:
			t := "FRONT_OF_CABIN"
			rowDesignation = &t
		case res.LastRowNumber:
			t := "VACANT_OR_OFFERED_LAST"
			rowDesignation = &t
		}

		var rowType *string
		switch {
		case rowNum == 10:
			t := "WING_BEGIN"
			rowType = &t
		case rowNum == 19:
			t := "WING_END"
			rowType = &t
		case rowNum >= 12 && rowNum <= 15:
			t := "EXIT"
			rowType = &t
		}

		// Create seat details
		var seats []entity.SeatDetail
		for _, col := range res.CabinLayouts {
			isNonSelectable := col.ColumnCode == "LEFT_SIDE" || col.ColumnCode == "RIGHT_SIDE" || col.ColumnCode == "AISLE"
			slotCode := "SEAT"
			if isNonSelectable {
				if isWing {
					slotCode = "WING"
				} else {
					slotCode = "BLANK"
				}
			}

			seats = append(seats, entity.SeatDetail{
				ID:                 uuid.New(),
				CabinID:            res.ID,
				CabinLayoutID:      col.ID,
				RowID:              rowID,
				StoreFrontSlotCode: slotCode,
				IsSelectable:       !isNonSelectable,
			})
		}

		cabinRowDetail := entity.CabinRow{
			ID:             rowID,
			CabinID:        res.ID,
			RowNumber:      rowNum,
			IsEnable:       true,
			RowType:        rowType,
			RowDesignation: rowDesignation,
			Seats:          seats, // Uncomment if used
		}

		if _, err := json.MarshalIndent(cabinRowDetail, "", "  "); err == nil {
			// fmt.Println(string(jsonData))
		}

		// repo.SetupCabinRow(context.Background(), &cabinRowDetail)
	}

}

func TestXxx(t *testing.T) {
	repo.SetupSeatsPrices(context.Background())
}

func TestYYYY(t *testing.T) {
	res, err := repo.GetSeatingWithPrice(context.Background(), "MYR", "69bf771a-b203-43da-b179-7c0ac3039386", )
	if err != nil {
		panic(err)
	}

	if jsonData, err := json.MarshalIndent(res, "", "  "); err == nil {
		fmt.Println(string(jsonData))
	}
}
