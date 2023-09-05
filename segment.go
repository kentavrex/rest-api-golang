package app

import "errors"

type Segment struct {
	Id   int    `json:"id" db:"id"`
	Slug string `json:"slug" db:"slug" binding:"required"`
}

type CreateSegmentInput struct {
	Slug string `json:"slug" db:"slug" binding:"required"`
}

type UpdateSegmentInput struct {
	Slug *string `json:"slug" db:"slug" binding:"required"`
}

func (i UpdateSegmentInput) Validate() error {
	if i.Slug == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
