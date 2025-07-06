package dto

type SeatView struct {
	StoreFrontSlotCode string
	IsSelectable bool
	FeeWaived bool
	FeeWaiverRuleID string
	RefundIndicator string
	FreeOfCharge bool
	IsAvailable bool
	Currency string
	PriceAmount float64
	TaxAmount float64
	Total float64
	RowNumber int
	IsEnable bool
	DisableCause string
	RowType string
	RowDesignation string
	ColumnCode string
}