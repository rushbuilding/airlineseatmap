package entity

type SpecialRequest struct {
	ID                 string `gorm:"column:special_request_id;primaryKey" json:"id"`
	SegmentPassengerID string `gorm:"column:segment_passenger_id" json:"segment_passenger_id"`
	RequestType        string `gorm:"column:request_type" json:"request_type"` // ENUM
	Description        string `gorm:"column:description" json:"description"`
}

func (s *SpecialRequest) TableName() string {
	return "special_request"
}