package entity

import "github.com/google/uuid"

type CabinConfiguration struct {
	ID           uuid.UUID `gorm:"column:cabin_layout_id;primaryKey" json:"id"`
	CabinID      uuid.UUID `gorm:"column:cabin_id" json:"cabin_id"`
	ColumnCode   string `gorm:"column:column_code" json:"column_code"`
	ColumnNumber int    `gorm:"column:column_number" json:"column_number"`
}

func (a *CabinConfiguration) TableName() string {
	return "cabin_configuration"
}