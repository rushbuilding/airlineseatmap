package controller

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/service"
)

type SeatControllerImpl struct {
	CabinService service.CabinService
	PassengerService service.PassengerService
	SegmentService service.SegmentService

}

func NewSeatController(cabinService service.CabinService, passengerService service.PassengerService, segmentService service.SegmentService) SeatingController {
	return &SeatControllerImpl{
		CabinService: cabinService,
		PassengerService: passengerService,
		SegmentService: segmentService,
	}
}

func (r *SeatControllerImpl) GetPassangerSeating(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := context.Background()
	passenger, _ := r.PassengerService.GetPassengerData(ctx, "3937094243023851424", "8fe24407-1df6-4b99-96c9-4dee6b72fdfb")
	segment, _ := r.SegmentService.GetSegmentData(ctx, "3937094243023851424", "ECONOMY")
	seating, _ := r.CabinService.GetSeatingMapAndPriceByCabinId(ctx, "3937094243023851424", "ECONOMY", "MYR")
}

func (r *SeatControllerImpl) SetUpData() {

}