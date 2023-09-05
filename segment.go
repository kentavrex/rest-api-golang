package app

import "errors"

type Segment struct {
	Id   int
	Slug string
}

type UpdateSegmentInput struct {
	Slug *string `json:"slug" db:"slug"`
}

func (i UpdateSegmentInput) Validate() error {
	if i.Slug == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
