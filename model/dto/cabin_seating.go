package dto

type PassengerSeatMap struct {
	SeatSelectionEnabledForPax bool
	SeatMap SeatMap
}

type SeatMap struct {
	RowsDisabledCauses []string
	Aircraft string
	Cabins []Cabin
}

type Cabin struct {
	Deck string
	SeatColumns []string
	SeatRows []SeatRow
}

type SeatRow struct {
	RowNumber int
	SeatCodes []string
	Seats []Seating
}

type Seating struct {
	StoreFrontSlotCode string
	Available bool
	Code string
	Designation []string
	Entitle bool
	FeeWaived bool
	EntitledRuleId string
	SeatCharacteristics []string
	Limitations []string
	RefundIndicator string
	FreeOfCharge bool
	Prices Price
	Taxes Price
	Total Price
	OriginallySelected bool
	RawSeatCharacteristics []string
}

type Price struct {
	Alternatives [][]PriceItem
}

type PriceItem struct {
	Amount float64
	Currency string
}