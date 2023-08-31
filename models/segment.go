package models

type Segment struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Slug string `json:"slug"`
}
