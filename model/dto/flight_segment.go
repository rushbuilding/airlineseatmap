package dto

type FlightSegment struct {
	Type string `json:"@Type"`
	SegmentOfferInformation SegmentOfferInformation
	Duration int
	CabinClass string
	Equipment string
	Flight FlightInfo
	Origin string
	Destination string
	Departure string
	Arrival string
	BookingClass string
	LayOverDuration string
	FareBasis string
	SubjectToGovernmentApproval bool
	SegmentRef string
}

type FlightInfo struct {
	FlightNumber int
	OperatingFlightNumber int
	AirlineCode string
	OperatingAirlineCode string
	StopsAirports []string
	DepartureTerminal string
	ArrivalTerminal string
}

type SegmentOfferInformation struct {
	FlightsMiles int
	AwardFare bool
}