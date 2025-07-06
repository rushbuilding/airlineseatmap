package entity

import "github.com/google/uuid"

type SeatDetail struct {
	ID                  uuid.UUID `gorm:"column:seat_id;primaryKey" json:"id"`
	CabinID             uuid.UUID `gorm:"column:cabin_id" json:"cabin_id"`
	RowID               uuid.UUID `gorm:"column:row_id" json:"row_id"`
	CabinLayoutID       uuid.UUID `gorm:"column:cabin_layout_id" json:"cabin_layout_id"`
	StoreFrontSlotCode  string `gorm:"column:store_front_slot_code" json:"store_front_slot_code"`
	IsSelectable        bool   `gorm:"column:is_selectable" json:"is_selectable"`
}

func (a *SeatDetail) TableName() string {
	return "seat_detail"
}