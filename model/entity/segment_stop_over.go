package entity

import "github.com/google/uuid"

type SegmentStopOver struct {
	StopOverID 		uuid.UUID 	`gorm:"column:stop_over_id;primarykey"`
	SegmentRef 		string 		`gorm:"column:segment_ref"`
	AirportCode 	string 		`gorm:"column:airport_code"`
	Duration 		uint8 		`gorm:"column:duration"`
}

func (f *SegmentStopOver) TableName() string {
	return "segment_stop_over"
}