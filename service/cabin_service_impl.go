package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/repository"
)

type CabinServiceImpl struct {
	AircraftRepository repository.AircraftManagementRepository
	SegmentRepository repository.SegmentManagementRepository
}

func NewCabinService(aircraftRepository repository.AircraftManagementRepository,segmentRepository repository.SegmentManagementRepository) CabinService {
	return &CabinServiceImpl{AircraftRepository:aircraftRepository, SegmentRepository: segmentRepository}
}

func (r *CabinServiceImpl) GetSeatingMapAndPriceByCabinId(ctx context.Context, segmentRef string, serviceClass string, currency string) (dto.SeatMap, error) {
	segment, _ := r.SegmentRepository.GetSegmentBySegmentRef(ctx, segmentRef)
	aircraft, err := r.AircraftRepository.GetAircraftByID(ctx, segment.AircraftID.String())
	if err != nil {
		panic(err)
	}

	cabin, err := r.AircraftRepository.GetCabinByAircraftIDAndServiceClass(ctx, segment.AircraftID.String(), serviceClass)
	if err != nil {
		panic(err)
	}

	res, err := r.AircraftRepository.GetSeatingWithPrice(ctx, currency, cabin.ID.String())

	if err != nil {
		panic(err)
	}

	var columns []string
	for _, c := range cabin.CabinLayouts {
		columns = append(columns, c.ColumnCode)
	}


	disableCauseSet := make(map[string]struct{})
	var disableCauses []string

	var seatRows []dto.SeatRow
	for i := 0; i < len(res); i += 9 {
		end := i + 9
		if end > len(res) {
			end = len(res)
		}

		rowSeats := res[i:end]
		if len(rowSeats) == 0 {
			continue
		}

		codeSet := make(map[string]struct{})
		var codes []string

		rowNum := rowSeats[0].RowNumber
		var seats []dto.Seating
		for _, seat := range rowSeats {
			if _, exists := codeSet[seat.StoreFrontSlotCode]; !exists {
				codeSet[seat.StoreFrontSlotCode] = struct{}{}
				codes = append(codes, seat.StoreFrontSlotCode)
			}

			if seat.DisableCause != "" {
				if _, exists := disableCauseSet[seat.DisableCause]; !exists {
					disableCauseSet[seat.DisableCause] = struct{}{}
					disableCauses = append(disableCauses, seat.DisableCause)
				}
	
			}

			var price dto.Price
			var tax dto.Price
			var total dto.Price
			var refundIndicator string
			var code string
			
			seating := dto.Seating{
				StoreFrontSlotCode: seat.StoreFrontSlotCode,
				Available: false,
				Entitle: false,
				FeeWaived: false,
				FreeOfCharge: true,
				OriginallySelected: false,
			}

			if seat.StoreFrontSlotCode == "SEAT" {
				code = fmt.Sprintf("%d%s", seat.RowNumber, seat.ColumnCode)
				refundIndicator = seat.RefundIndicator
				price = dto.Price{
					Alternatives: [][]dto.PriceItem{
						{
							{Amount: seat.PriceAmount, Currency: seat.Currency},
						},
					},
				}
				tax = dto.Price{
					Alternatives: [][]dto.PriceItem{
						{
							{Amount: seat.TaxAmount, Currency: seat.Currency},
						},
					},
				}
				total = dto.Price{
					Alternatives: [][]dto.PriceItem{
						{
							{Amount: seat.Total, Currency: seat.Currency},
						},
					},
				}

				seating = dto.Seating{
					StoreFrontSlotCode: seat.StoreFrontSlotCode,
					Available: seat.IsAvailable,
					Code: code,
					Entitle: true,
					FeeWaived: seat.FeeWaived,
					EntitledRuleId: "",
					RefundIndicator: refundIndicator,
					FreeOfCharge: seat.FreeOfCharge,
					Prices: price,
					Taxes: tax,
					Total: total,
					OriginallySelected: false,
				}
			}
			
			seats = append(seats, seating)
		}

		row := dto.SeatRow{
			RowNumber: rowNum,
			SeatCodes: codes,
			Seats: seats,
		}

		seatRows = append(seatRows, row)
	}

	cabinDeck := dto.Cabin{
		Deck: cabin.Deck,
		SeatColumns: columns,
		SeatRows: seatRows,
	}

	passengerSeatMap := dto.SeatMap{
		RowsDisabledCauses: disableCauses,
		Aircraft:  strconv.Itoa(int(aircraft.RegistrationNumber)),
		Cabins: []dto.Cabin{cabinDeck},
	}

	return passengerSeatMap, nil
}