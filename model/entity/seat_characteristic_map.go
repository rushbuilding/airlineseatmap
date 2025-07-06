package entity

type SeatCharacteristicMap struct {
	SeatID          string `gorm:"column:seat_id;primaryKey" json:"seat_id"`
	CharacteristicID string `gorm:"column:characteristic_id;primaryKey" json:"characteristic_id"`
	IsDisplay       bool   `gorm:"column:is_display" json:"is_display"`

	// Associations
	Characteristic SeatCharacteristic `gorm:"foreignKey:characteristic_id;references:characteristic_id" json:"characteristic"`
}

func (a *SeatCharacteristicMap) TableName() string {
	return "seat_characteristic_map"
}