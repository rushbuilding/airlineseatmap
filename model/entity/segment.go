package entity

import "time"

type Segment struct {
	SegmentRef 					string 			`gorm:"column:segment_ref;primarykey"`
	OriginAirportCode 			int16 			`gorm:"column:flight_number"`
	DestinationAirportCode 		int16 			`gorm:"column:operating_flight_number"`
	DepartureTime 				time.Time 		`gorm:"column:departure_time"`
	DepatureTimezone			string 			`gorm:"column:departure_timezone"`
	ArrivalTime 				time.Time 		`gorm:"column:arrival_time"`
	ArrivalTimezone				string 			`gorm:"column:arival_timezone"`
	Duration					int16			`gorm:"column:duration"`
	DepartureTerminal			string			`gorm:"column:departure_terminal"`
	ArrivalTerminal				string			`gorm:"column:arrival_terminal"`
	SubjectGovermentApproval 	bool 			`gorm:"column:subject_goverment_approval"`
}

func (f *Segment) TableName() string {
	return "segment"
}