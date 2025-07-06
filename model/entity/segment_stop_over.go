package entity

import "github.com/google/uuid"

type SegmentStopOver struct {
	ID          uuid.UUID `gorm:"column:stop_over_id;primaryKey" json:"id"`
	SegmentRef  string    `gorm:"column:segment_ref" json:"segment_ref"`
	AirportCode string    `gorm:"column:airport_code" json:"airport_code"`
	Duration    uint8     `gorm:"column:duration" json:"duration"`
}

func (a *SegmentStopOver) TableName() string {
	return "segment_stop_over"
}