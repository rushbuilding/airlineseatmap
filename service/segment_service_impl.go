package service

import (
	"context"
	"strconv"

	"github.com/rushbuilding/airlineseatmap/model/dto"
	"github.com/rushbuilding/airlineseatmap/repository"
)

type SegmentServiceImpl struct {
	SegmentRepository repository.SegmentManagementRepository
}

func NewSegmentService(SegmentRepository repository.SegmentManagementRepository) SegmentService {
	return &SegmentServiceImpl{SegmentRepository:SegmentRepository}
}

func (r *SegmentServiceImpl) GetSegmentData(ctx context.Context, ref string, class string) (*dto.FlightSegment, error) {
	segmentOffer, err := r.SegmentRepository.GetSegmentOfferingBySegmentRefAndServiceClass(ctx, ref, class)

	if err != nil {
		panic(err)
	}
	segment, _ := r.SegmentRepository.GetSegmentBySegmentRef(ctx, ref)

	offer := dto.SegmentOfferInformation{
		FlightsMiles: segmentOffer.FlightMiles,
		AwardFare: segmentOffer.AwardFare,
	}

	var stopAirporCode []string
	var layover int64 = 0
	for _, p := range segment.StopOvers {
		stopAirporCode = append(stopAirporCode, p.AirportCode)
		layover += int64(p.Duration)
	}

	flightInfo := dto.FlightInfo{
		FlightNumber: int(segment.FlightDetail.FlightNumber),
		OperatingFlightNumber: int(segment.FlightDetail.OperatingFlightNumber),
		AirlineCode: segment.FlightDetail.AirlineCode,
		OperatingAirlineCode: segment.FlightDetail.OperatingAirlineCode,
		StopsAirports: stopAirporCode,
		DepartureTerminal: segment.DepartureTerminal,
		ArrivalTerminal: segment.ArrivalTerminal,
	}

	flightDetail := dto.FlightSegment{
		Type: "Segment",
		SegmentOfferInformation: offer,
		Duration: int(segment.Duration),
		CabinClass: segmentOffer.ServiceClass,
		Equipment: strconv.Itoa(int(segment.Aircraft.RegistrationNumber)),
		Flight: flightInfo,
		Origin: segment.OriginAirportCode,
		Destination: segment.DestinationAirportCode,
		Departure: segment.DepartureTime.Format(""),
		Arrival: segment.ArrivalTime.Format(""),
		BookingClass: segmentOffer.BookingClass,
		LayOverDuration:  strconv.Itoa(int(layover)),
		FareBasis: segmentOffer.FareBasis,
		SubjectToGovernmentApproval: segment.SubjectGovermentApproval,
		SegmentRef: segment.SegmentRef,
	}

	return &flightDetail, nil

}