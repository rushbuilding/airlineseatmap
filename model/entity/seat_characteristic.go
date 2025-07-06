package entity

type SeatCharacteristic struct {
	ID             string `gorm:"column:characteristic_id;primaryKey" json:"id"`
	LimitationCode string `gorm:"column:limitation_code" json:"limitation_code"`
}

func (a *SeatCharacteristic) TableName() string {
	return "seat_characteristic"
}