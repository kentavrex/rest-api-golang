package models

type UserSegment struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UserId    int
	SegmentId int
}
